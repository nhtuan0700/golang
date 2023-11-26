module example.local/hello

go 1.21.1

replace example.local/greetings => ../greetings

require example.local/greetings v0.0.0-00010101000000-000000000000
