# iroiro
tl;dr: I created it for a LINE chatbot that converts various units. I was tired of Googling and trying to convert values in my head whenever talking to Japanese friends so I put all the kinds of conversions I need to do in one place.

It's probably not very useful to someone who's not me, but I wanted it as a module so I could use it across a couple repos.

Types of conversions:
- Japanese numbers <-> decimal: e.g. 10,000 => 10万
- JPY <-> USD
- imperial <-> metric (right now, only distance)

# Functionality

`amountconvert` does the overall functionality called by the LINE chatbot but the other packages can be used.

## `amountconverter` package
Utilities for taking input and parsing whether it's only a number, a currency amount, or a measurement. Uses functions from `counting`, `currencyconverter`, `distanceconverter` and `stringutils` to both convert the values, and make them readable. 

## `counting` package
Number parsing and conversion to JP units. Function to parse string in various formats (e.g. 135k, 1万) and convert to a float, and function to convert to Japanese number. Uses `stringutils` for prettifying some outputs.

## `currencyconverter` package
Utilities to convert from JPY <-> USD. It uses a rate passed in by caller.

## `distanceconverter` package
Utilities to convert imperial to metric and vice versa. Uses `stringutils` for prettifying the output.

## `stringutils` package
Utilities for trimming floats, and for adding commas to big numbers.