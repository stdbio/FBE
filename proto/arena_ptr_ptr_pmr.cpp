//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: arena_ptr.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

#include "arena_ptr_ptr_pmr.h"

namespace arena_ptr_pmr {

Line::Line()
    : expression()
{}

Line::Line([[maybe_unused]] allocator_type alloc)
    : expression(assign_member<::arena_common_pmr::Expression>(alloc))
{}

Line::Line(::arena_common_pmr::Expression arg_expression)
    : expression(std::move(arg_expression))
{}

Line::Line([[maybe_unused]] Line&& other) noexcept
    : expression(std::move(other.expression))
{}

Line::~Line()
{
}

bool Line::operator==([[maybe_unused]] const Line& other) const noexcept
{
    if (expression != other.expression)
        return false;
    return true;
}

bool Line::operator<([[maybe_unused]] const Line& other) const noexcept
{
    return false;
}

Line& Line::operator=(Line&& other) noexcept
{
    if (this != &other)
    {
        expression = std::move(other.expression);
    }
    return *this;
}

std::string Line::string() const
{
    std::stringstream ss; ss << *this; return ss.str();
}

void Line::swap([[maybe_unused]] Line& other) noexcept
{
    using std::swap;
    swap(expression, other.expression);
}

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Line& value)
{
    stream << "Line(";
    stream << "expression="; stream << value.expression;
    stream << ")";
    return stream;
}

Line2::Line2()
    : bytes_v()
{}

Line2::Line2([[maybe_unused]] allocator_type alloc)
    : bytes_v(alloc)
{}

Line2::Line2(const FBE::pmr_buffer_t& arg_bytes_v)
    : bytes_v(arg_bytes_v)
{}

Line2::Line2([[maybe_unused]] Line2&& other) noexcept
    : bytes_v(std::move(other.bytes_v))
{}

Line2::~Line2()
{
}

bool Line2::operator==([[maybe_unused]] const Line2& other) const noexcept
{
    if (bytes_v != other.bytes_v)
        return false;
    return true;
}

bool Line2::operator<([[maybe_unused]] const Line2& other) const noexcept
{
    return false;
}

Line2& Line2::operator=(Line2&& other) noexcept
{
    if (this != &other)
    {
        bytes_v = std::move(other.bytes_v);
    }
    return *this;
}

std::string Line2::string() const
{
    std::stringstream ss; ss << *this; return ss.str();
}

void Line2::swap([[maybe_unused]] Line2& other) noexcept
{
    using std::swap;
    swap(bytes_v, other.bytes_v);
}

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Line2& value)
{
    stream << "Line2(";
    stream << "bytes_v="; stream << "bytes[" << value.bytes_v.size() << "]";
    stream << ")";
    return stream;
}

} // namespace arena_ptr_pmr
