//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

#pragma once

#ifdef isset
#undef isset
#endif

#if defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#elif defined(_MSC_VER)
#pragma system_header
#endif

#include "fbe.h"

namespace proto {
using namespace FBE;
} // namespace proto

namespace FBE {
using namespace ::proto;
} // namespace FBE

namespace proto {

enum class OrderSide : uint8_t
{
    buy,
    sell,
};

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] OrderSide value);

#if defined(FMT_VERSION)
} template <> struct fmt::formatter<proto::OrderSide> : formatter<string_view> {}; namespace proto {
#endif

#if defined(LOGGING_PROTOCOL)
CppLogging::Record& operator<<(CppLogging::Record& record, OrderSide value);
#endif

enum class OrderType : uint8_t
{
    market,
    limit,
    stop,
};

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] OrderType value);

#if defined(FMT_VERSION)
} template <> struct fmt::formatter<proto::OrderType> : formatter<string_view> {}; namespace proto {
#endif

#if defined(LOGGING_PROTOCOL)
CppLogging::Record& operator<<(CppLogging::Record& record, OrderType value);
#endif

enum class State : uint8_t
{
    unknown = (uint8_t)0x00u,
    invalid = (uint8_t)0x01u,
    initialized = (uint8_t)0x02u,
    calculated = (uint8_t)0x04u,
    broken = (uint8_t)0x08u,
    good = initialized  |  calculated,
    bad = unknown  |  invalid  |  broken,
};

FBE_ENUM_FLAGS(State)

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] State value);

#if defined(FMT_VERSION)
} template <> struct fmt::formatter<proto::State> : formatter<string_view> {}; namespace proto {
#endif

#if defined(LOGGING_PROTOCOL)
CppLogging::Record& operator<<(CppLogging::Record& record, State value);
#endif

struct Order
{
    int32_t id;
    stdb::memory::string symbol;
    ::proto::OrderSide side;
    ::proto::OrderType type;
    double price;
    double volume;

    size_t fbe_type() const noexcept { return 1; }

    Order();
    Order(int32_t arg_id, const stdb::memory::string& arg_symbol, const ::proto::OrderSide& arg_side, const ::proto::OrderType& arg_type, double arg_price, double arg_volume);
    Order(const Order& other) = default;
    Order(Order&& other) = default;
    ~Order() = default;

    Order& operator=(const Order& other) = default;
    Order& operator=(Order&& other) = default;

    bool operator==(const Order& other) const noexcept;
    bool operator!=(const Order& other) const noexcept { return !operator==(other); }
    bool operator<(const Order& other) const noexcept;
    bool operator<=(const Order& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Order& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Order& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Order& value);

    void swap(Order& other) noexcept;
    friend void swap(Order& value1, Order& value2) noexcept { value1.swap(value2); }
};

} // namespace proto

#if defined(FMT_VERSION)
template <> struct fmt::formatter<proto::Order> : formatter<string_view> {};
#endif

template<>
struct std::hash<proto::Order>
{
    typedef proto::Order argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        result = result * 31 + std::hash<decltype(value.id)>()(value.id);
        return result;
    }
};

namespace proto {

struct Balance
{
    stdb::memory::string currency;
    double amount;

    size_t fbe_type() const noexcept { return 2; }

    Balance();
    Balance(const stdb::memory::string& arg_currency, double arg_amount);
    Balance(const Balance& other) = default;
    Balance(Balance&& other) = default;
    ~Balance() = default;

    Balance& operator=(const Balance& other) = default;
    Balance& operator=(Balance&& other) = default;

    bool operator==(const Balance& other) const noexcept;
    bool operator!=(const Balance& other) const noexcept { return !operator==(other); }
    bool operator<(const Balance& other) const noexcept;
    bool operator<=(const Balance& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Balance& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Balance& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Balance& value);

    void swap(Balance& other) noexcept;
    friend void swap(Balance& value1, Balance& value2) noexcept { value1.swap(value2); }
};

} // namespace proto

#if defined(FMT_VERSION)
template <> struct fmt::formatter<proto::Balance> : formatter<string_view> {};
#endif

template<>
struct std::hash<proto::Balance>
{
    typedef proto::Balance argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        result = result * 31 + std::hash<decltype(value.currency)>()(value.currency);
        return result;
    }
};

namespace proto {

struct Account
{
    int32_t id;
    stdb::memory::string name;
    ::proto::State state;
    ::proto::Balance wallet;
    std::optional<::proto::Balance> asset;
    std::vector<::proto::Order> orders;

    size_t fbe_type() const noexcept { return 3; }

    Account();
    Account(int32_t arg_id, const stdb::memory::string& arg_name, const ::proto::State& arg_state, const ::proto::Balance& arg_wallet, const std::optional<::proto::Balance>& arg_asset, const std::vector<::proto::Order>& arg_orders);
    Account(const Account& other) = default;
    Account(Account&& other) = default;
    ~Account() = default;

    Account& operator=(const Account& other) = default;
    Account& operator=(Account&& other) = default;

    bool operator==(const Account& other) const noexcept;
    bool operator!=(const Account& other) const noexcept { return !operator==(other); }
    bool operator<(const Account& other) const noexcept;
    bool operator<=(const Account& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Account& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Account& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Account& value);

    void swap(Account& other) noexcept;
    friend void swap(Account& value1, Account& value2) noexcept { value1.swap(value2); }
};

} // namespace proto

#if defined(FMT_VERSION)
template <> struct fmt::formatter<proto::Account> : formatter<string_view> {};
#endif

template<>
struct std::hash<proto::Account>
{
    typedef proto::Account argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        result = result * 31 + std::hash<decltype(value.id)>()(value.id);
        return result;
    }
};

namespace proto {

struct CharMap
{
    std::unordered_map<char, stdb::memory::string> abbr;

    size_t fbe_type() const noexcept { return 1; }

    CharMap();
    explicit CharMap(const std::unordered_map<char, stdb::memory::string>& arg_abbr);
    CharMap(const CharMap& other) = default;
    CharMap(CharMap&& other) = default;
    ~CharMap() = default;

    CharMap& operator=(const CharMap& other) = default;
    CharMap& operator=(CharMap&& other) = default;

    bool operator==(const CharMap& other) const noexcept;
    bool operator!=(const CharMap& other) const noexcept { return !operator==(other); }
    bool operator<(const CharMap& other) const noexcept;
    bool operator<=(const CharMap& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const CharMap& other) const noexcept { return !operator<=(other); }
    bool operator>=(const CharMap& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const CharMap& value);

    void swap(CharMap& other) noexcept;
    friend void swap(CharMap& value1, CharMap& value2) noexcept { value1.swap(value2); }
};

} // namespace proto

#if defined(FMT_VERSION)
template <> struct fmt::formatter<proto::CharMap> : formatter<string_view> {};
#endif

template<>
struct std::hash<proto::CharMap>
{
    typedef proto::CharMap argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace proto {

struct OrderMessage
{
    ::proto::Order body;

    size_t fbe_type() const noexcept { return 2; }

    OrderMessage();
    explicit OrderMessage(const ::proto::Order& arg_body);
    OrderMessage(const OrderMessage& other) = default;
    OrderMessage(OrderMessage&& other) = default;
    ~OrderMessage() = default;

    OrderMessage& operator=(const OrderMessage& other) = default;
    OrderMessage& operator=(OrderMessage&& other) = default;

    bool operator==(const OrderMessage& other) const noexcept;
    bool operator!=(const OrderMessage& other) const noexcept { return !operator==(other); }
    bool operator<(const OrderMessage& other) const noexcept;
    bool operator<=(const OrderMessage& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const OrderMessage& other) const noexcept { return !operator<=(other); }
    bool operator>=(const OrderMessage& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const OrderMessage& value);

    void swap(OrderMessage& other) noexcept;
    friend void swap(OrderMessage& value1, OrderMessage& value2) noexcept { value1.swap(value2); }
};

} // namespace proto

#if defined(FMT_VERSION)
template <> struct fmt::formatter<proto::OrderMessage> : formatter<string_view> {};
#endif

template<>
struct std::hash<proto::OrderMessage>
{
    typedef proto::OrderMessage argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace proto {

struct BalanceMessage
{
    ::proto::Balance body;

    size_t fbe_type() const noexcept { return 3; }

    BalanceMessage();
    explicit BalanceMessage(const ::proto::Balance& arg_body);
    BalanceMessage(const BalanceMessage& other) = default;
    BalanceMessage(BalanceMessage&& other) = default;
    ~BalanceMessage() = default;

    BalanceMessage& operator=(const BalanceMessage& other) = default;
    BalanceMessage& operator=(BalanceMessage&& other) = default;

    bool operator==(const BalanceMessage& other) const noexcept;
    bool operator!=(const BalanceMessage& other) const noexcept { return !operator==(other); }
    bool operator<(const BalanceMessage& other) const noexcept;
    bool operator<=(const BalanceMessage& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const BalanceMessage& other) const noexcept { return !operator<=(other); }
    bool operator>=(const BalanceMessage& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const BalanceMessage& value);

    void swap(BalanceMessage& other) noexcept;
    friend void swap(BalanceMessage& value1, BalanceMessage& value2) noexcept { value1.swap(value2); }
};

} // namespace proto

#if defined(FMT_VERSION)
template <> struct fmt::formatter<proto::BalanceMessage> : formatter<string_view> {};
#endif

template<>
struct std::hash<proto::BalanceMessage>
{
    typedef proto::BalanceMessage argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace proto {

struct AccountMessage
{
    ::proto::Account body;

    size_t fbe_type() const noexcept { return 4; }

    AccountMessage();
    explicit AccountMessage(const ::proto::Account& arg_body);
    AccountMessage(const AccountMessage& other) = default;
    AccountMessage(AccountMessage&& other) = default;
    ~AccountMessage() = default;

    AccountMessage& operator=(const AccountMessage& other) = default;
    AccountMessage& operator=(AccountMessage&& other) = default;

    bool operator==(const AccountMessage& other) const noexcept;
    bool operator!=(const AccountMessage& other) const noexcept { return !operator==(other); }
    bool operator<(const AccountMessage& other) const noexcept;
    bool operator<=(const AccountMessage& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const AccountMessage& other) const noexcept { return !operator<=(other); }
    bool operator>=(const AccountMessage& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const AccountMessage& value);

    void swap(AccountMessage& other) noexcept;
    friend void swap(AccountMessage& value1, AccountMessage& value2) noexcept { value1.swap(value2); }
};

} // namespace proto

#if defined(FMT_VERSION)
template <> struct fmt::formatter<proto::AccountMessage> : formatter<string_view> {};
#endif

template<>
struct std::hash<proto::AccountMessage>
{
    typedef proto::AccountMessage argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace proto {

} // namespace proto
