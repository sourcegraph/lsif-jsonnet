local common = import './common.libsonnet';

// Define a local function.
// Default arguments are like Python:
local foo(x, y=10) = x + y;

local my_value = {
  // A method
  foo(x): x * x,
};

{
  // Functions are first class citizens.
  call_inline_function:
    (function(x) x * x)(5),

  // Using the variable fetches the function,
  // the parens call the function.
  call: foo(2),

  common_call: common.bar(3, 2),

  // Like python, parameters can be named at
  // call time.
  named_params: foo(x=2),
  // This allows changing their order
  named_params2: foo(y=3, x=2),

  // my_value.foo returns the function,
  // which is then called like any other.
  call_method1: my_value.foo(3),

  standard_lib:
    std.join(' ', std.split('foo/bar', '/')),
  len: [
    std.length('hello'),
    std.length([1, 2, 3]),
  ],
}
