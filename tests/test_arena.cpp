/*
* Copyright (C) 2020 Beijing Jinyi Data Technology Co., Ltd. All rights reserved.
*
* This source code is licensed under the MIT license found in the
* LICENSE file in the root directory of this source tree.
*/
#include <string>

#include "arena_pmr.h"
#include "arena_ptr_ptr_pmr.h"
#include "catch2/catch.hpp"
#include "fbe.h"
#include "test.h"
#include "arena/arena.hpp"
#include "../proto/arena_pmr_models.h"
#include "../proto/arena_ptr_ptr_pmr_models.h"

using stdb::memory::Arena;
using stdb::memory::ArenaContainStatus;

class alloc_class
{
   public:
    alloc_class() : init(static_cast<char*>(malloc(4 * 1024 * 1024))), curr(init) {}
    virtual ~alloc_class() { free(init); }

    virtual auto alloc(uint64_t size) -> void* {
        alloc_sizes.push_back(size);
        void* ret = static_cast<void*>(curr);
        ptrs.push_back(ret);

        curr += size;
        return ret;
    }

    void dealloc(void* ptr) { free_ptrs.push_back(ptr); }

    void reset() {
        alloc_sizes.clear();
        ptrs.clear();
        free_ptrs.clear();
    }

    // private:
    std::vector<uint64_t> alloc_sizes;
    std::vector<void*> ptrs;
    std::vector<void*> free_ptrs;
    char* init{nullptr};
    char* curr{nullptr};
};

class ArenaTest
{
   public:
    Arena::Options ops_simple;

    thread_local static alloc_class* mock;

    static auto mock_alloc(std::size_t size) -> void* { return mock->alloc(size); }

    static void mock_dealloc(void* ptr) { return mock->dealloc(ptr); }

    ArenaTest() {
        // initialize the ops_simple
        ops_simple.block_alloc = &mock_alloc;
        ops_simple.block_dealloc = &mock_dealloc;
        ops_simple.normal_block_size = 1024ULL;
        ops_simple.suggested_init_block_size = 0ULL;
        ops_simple.huge_block_size = 0ULL;
    }

    template <typename Func>
    auto test_fn_with_allocs(Func&& func) {
        mock = new alloc_class;
        auto* arena = new Arena(ops_simple);

        REQUIRE(mock->ptrs.size() == 0);
        REQUIRE(mock->free_ptrs.size() == 0);

        func(*arena);

        delete arena;
        REQUIRE(mock->ptrs.size() != 0);
        REQUIRE(mock->free_ptrs.size() != 0);
        delete mock;
    }
};

thread_local alloc_class* ArenaTest::mock = nullptr;

// 1. template arena simple test
// 2. ptr arena simple test
// 2. ptr import template cstr
// 3. template import template cstr
// 4. ptr import ptr import

// bytes
// vector 
// map,
// struct
// variant and so on

TEST_CASE_METHOD(ArenaTest, "Arena (template import template)", "[template-based FBE]") {
  test_fn_with_allocs([](Arena& arena) {
    ::arena_common_pmr::Expression expr(arena.get_memory_resource());
    expr.keys.reserve(1);
    expr.aliases.reserve(1);

    FBE::ArenaString name("an alias name", arena.get_memory_resource());
    expr.keys.emplace_back(std::move(name));

    ::arena_common_pmr::Alias alias(arena.get_memory_resource());
    alias.name.assign("an alias name 2");
    alias.optr = ::arena_common_pmr::Optr::GE;
    expr.aliases.emplace_back(std::move(alias));

    ::arena_common_pmr::Alias alias2(arena.get_memory_resource());
    alias2.name.assign("an alias name 3");
    alias2.optr = ::arena_common_pmr::Optr::GT;
    expr.alias_int.emplace(24, std::move(alias2));

    auto* item = arena.Create<::arena_pmr::Item>();
    item->optr = ::arena_common_pmr::Optr::LE;
    item->alias.name.assign("an alias name for item");
    item->alias.optr = ::arena_common_pmr::Optr::LT;
    item->alias.expr.emplace<int32_t>(42);
    item->expressions.emplace_back(std::move(expr));

    REQUIRE(arena.check(reinterpret_cast<char*>(&item->optr)) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(item->alias.name.data())) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(&item->alias.optr)) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(&item->alias.expr)) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(&item->expressions)) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(&item->aliases_int)) == ArenaContainStatus::BlockUsed);
    auto& item_expr = item->expressions.at(0);
    REQUIRE(arena.check(reinterpret_cast<char*>(item_expr.keys.at(0).data())) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(item_expr.aliases.at(0).name.data())) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(&item_expr.aliases.at(0).optr)) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(item_expr.alias_int.at(24).name.data())) == ArenaContainStatus::BlockUsed);
    REQUIRE(arena.check(reinterpret_cast<char*>(&item_expr.alias_int.at(24).optr)) == ArenaContainStatus::BlockUsed);

    ::arena_common_pmr::Alias alias4(arena.get_memory_resource());
    alias4.name.assign("an alias name 4");
    alias4.optr = ::arena_common_pmr::Optr::EQ;
    item->aliases_int.emplace(42, std::move(alias4));

    FBE::arena_pmr::ItemModel item_writer;
    size_t serialized_info = item_writer.serialize(*item, arena.get_memory_resource());
    REQUIRE(serialized_info == item_writer.buffer().size());
    REQUIRE(item_writer.verify());

    FBE::arena_pmr::ItemModel item_reader;
    item_reader.attach(item_writer.buffer());
    REQUIRE(item_reader.verify());

    ::arena_pmr::Item copy;
    size_t deserialized_info = item_reader.deserialize(copy, arena.get_memory_resource());
    REQUIRE(deserialized_info == item_reader.buffer().size());

    REQUIRE(copy.optr == ::arena_common_pmr::Optr::LE);
    REQUIRE(copy.alias.name == "an alias name for item");
    REQUIRE(copy.alias.optr == ::arena_common_pmr::Optr::LT);
    REQUIRE(copy.alias.expr.index() == 2);
    REQUIRE(std::get<2>(copy.alias.expr) == 42);
    REQUIRE(copy.aliases_int.size() == 1);
    REQUIRE(copy.aliases_int.at(42).name == "an alias name 4");
    REQUIRE(copy.aliases_int.at(42).optr == ::arena_common_pmr::Optr::EQ);

    REQUIRE(copy.expressions.size() == 1);
    auto& expression = copy.expressions.at(0);
    REQUIRE(expression.keys.size() == 1);
    REQUIRE(expression.keys.at(0) == "an alias name");
    REQUIRE(expression.aliases.size() == 1);
    REQUIRE(expression.aliases.at(0).name == "an alias name 2");
    REQUIRE(expression.aliases.at(0).optr == ::arena_common_pmr::Optr::GE);
    REQUIRE(expression.alias_int.size() == 1);
    REQUIRE(expression.alias_int.at(24).name == "an alias name 3");
    REQUIRE(expression.alias_int.at(24).optr == ::arena_common_pmr::Optr::GT);
    REQUIRE(*item == copy);
  });
}

TEST_CASE_METHOD(ArenaTest, "Arena (ptr import template)", "[ptr-based FBE]") {
    test_fn_with_allocs([](Arena& arena) {
        auto* line = arena.Create<::arena_ptr_pmr::Line>();

        line->expression.keys.reserve(1);
        line->expression.aliases.reserve(1);

        FBE::ArenaString name("an alias name", arena.get_memory_resource());
        line->expression.keys.emplace_back(std::move(name));

        ::arena_common_pmr::Alias alias(arena.get_memory_resource());
        alias.name.assign("an alias name 2");
        alias.optr = ::arena_common_pmr::Optr::GE;
        line->expression.aliases.emplace_back(std::move(alias));

        ::arena_common_pmr::Alias alias2(arena.get_memory_resource());
        alias2.name.assign("an alias name 3");
        alias2.optr = ::arena_common_pmr::Optr::GT;
        line->expression.alias_int.emplace(24, std::move(alias2));

        ::arena_common_pmr::Alias alias4(arena.get_memory_resource());
        alias4.name.assign("an alias name long enough to alloc on arena");
        alias4.optr = ::arena_common_pmr::Optr::EQ;
        line->expression.alias_int.emplace(42, std::move(alias4));

        REQUIRE(arena.check(reinterpret_cast<char*>(&line->expression)) == ArenaContainStatus::BlockUsed);
        auto& line_expr = line->expression;
        REQUIRE(arena.check(reinterpret_cast<char*>(line_expr.keys.at(0).data())) == ArenaContainStatus::BlockUsed);
        REQUIRE(arena.check(reinterpret_cast<char*>(line_expr.aliases.at(0).name.data())) == ArenaContainStatus::BlockUsed);
        REQUIRE(arena.check(reinterpret_cast<char*>(&line_expr.aliases.at(0).optr)) == ArenaContainStatus::BlockUsed);
        REQUIRE(arena.check(reinterpret_cast<char*>(line_expr.alias_int.at(24).name.data())) == ArenaContainStatus::BlockUsed);
        REQUIRE(arena.check(reinterpret_cast<char*>(&line_expr.alias_int.at(24).optr)) == ArenaContainStatus::BlockUsed);
        REQUIRE(arena.check(reinterpret_cast<char*>(line_expr.alias_int.at(42).name.data())) == ArenaContainStatus::BlockUsed);
        REQUIRE(arena.check(reinterpret_cast<char*>(&line_expr.alias_int.at(42).optr)) == ArenaContainStatus::BlockUsed);

        FBE::arena_ptr_pmr::LineModel line_writer;
        size_t serialized_info = line_writer.serialize(*line, arena.get_memory_resource());
        REQUIRE(serialized_info == line_writer.buffer().size());
        REQUIRE(line_writer.verify());

        FBE::arena_ptr_pmr::LineModel line_reader;
        line_reader.attach(line_writer.buffer());
        REQUIRE(line_reader.verify());

        ::arena_ptr_pmr::Line line_copy;
        size_t deserialized_info = line_reader.deserialize(line_copy, arena.get_memory_resource());
        REQUIRE(deserialized_info == line_reader.buffer().size());

        auto& expression = line_copy.expression;
        REQUIRE(expression.keys.size() == 1);
        REQUIRE(expression.keys.at(0) == "an alias name");
        REQUIRE(expression.aliases.size() == 1);
        REQUIRE(expression.aliases.at(0).name == "an alias name 2");
        REQUIRE(expression.aliases.at(0).optr == ::arena_common_pmr::Optr::GE);
        REQUIRE(expression.alias_int.size() == 2);
        REQUIRE(expression.alias_int.at(24).name == "an alias name 3");
        REQUIRE(expression.alias_int.at(24).optr == ::arena_common_pmr::Optr::GT);
        REQUIRE(expression.alias_int.at(42).name == "an alias name long enough to alloc on arena");
        REQUIRE(expression.alias_int.at(42).optr == ::arena_common_pmr::Optr::EQ);
        REQUIRE(*line == line_copy);
    });
}

TEST_CASE_METHOD(ArenaTest, "Arena (pmr::bytes)", "[ptr-based FBE]") {
    test_fn_with_allocs([](Arena& arena) {
        auto* line2 = arena.Create<::arena_ptr_pmr::Line2>();

        pmr::vector<uint8_t> bytes_v{{65, 66, 67, 68, 69}, arena.get_memory_resource()};
        FBE::pmr_buffer_t buffer(std::move(bytes_v));
        line2->bytes_v = std::move(buffer);

        REQUIRE(arena.check(reinterpret_cast<char*>(&line2->bytes_v)) == ArenaContainStatus::BlockUsed);
        REQUIRE(arena.check(reinterpret_cast<char*>(line2->bytes_v.data())) == ArenaContainStatus::BlockUsed);

        FBE::arena_ptr_pmr::Line2Model line_writer;
        size_t serialized_info = line_writer.serialize(*line2, arena.get_memory_resource());
        REQUIRE(serialized_info == line_writer.buffer().size());
        CHECK(line_writer.verify());

        FBE::arena_ptr_pmr::Line2Model line_reader;
        line_reader.attach(line_writer.buffer());
        CHECK(line_reader.verify());

        ::arena_ptr_pmr::Line2 copy_line;
        size_t deserialized_info = line_reader.deserialize(copy_line, arena.get_memory_resource());
        REQUIRE(deserialized_info == line_reader.buffer().size());

        REQUIRE(line2->bytes_v.string() == "ABCDE");
        REQUIRE(*line2 == copy_line);
    });
}
