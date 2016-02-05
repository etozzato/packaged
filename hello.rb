require 'ffi'

class Hello

  extend FFI::Library
  ffi_lib 'packaged.so'
  attach_function :wthello, [:string], :string
  attach_function :wchello, [:string], :string

end

hh = Hello.new

puts hh.wthello("mek")
puts hh.wchello("mek")
