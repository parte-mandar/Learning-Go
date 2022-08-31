module example.com/callGreetings

go 1.19

replace example.com/greetings => ../greetings // go mod edit -replace example.com/greetings=../greetings

// This was done to make sure that main module looks for the greetings module in right directiry

require example.com/greetings v0.0.0-00010101000000-000000000000 // go mod tidy
