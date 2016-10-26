# Discounted-Offers

This repo provides a solution to the code challenge [Disacount-Offers](https://www.codeeval.com/public_sc/48/ "Disacount-Offers")


## Building

You need a Go development environment with both `GOROOT` and `GOPATH`
variables set. To install all dependencies, change directory into the project and run:

	go get ./...

To build, run:

	go build

## Testing

To run unit test, run:

	go test -v -race

## Running

- To run the program, create an input file:

		Jack Abraham,John Evans,Ted Dziuba;iPad 2 - 4-pack,Girl Scouts Thin Mints,Nerf Crossbow

		Jeffery Lebowski,Walter Sobchak,Theodore Donald Kerabatsos,Peter Gibbons,Michael Bolton,Samir Nagheenanajar;Half & Half,Colt M1911A1,16lb bowling ball,Red Swingline Stapler,Printer paper,Vibe Magazine Subscriptions - 40 pack

		Jareau Wade,Rob Eroh,Mahmoud Abdelkader,Wenyi Cai,Justin Van Winkle,Gabriel Sinkin,Aaron Adelson;Batman No. 1,Football - Official Size,Bass Amplifying Headphones,Elephant food - 1024 lbs,Three Wolf One Moon T-shirt,Dom Perignon 2000 Vintage`

- Run the program with the input filename as the only parameter

		C:\dev\gocode\src\github.com\Discounted-Offers\input>discounted-offers exampleInput
		21.00
		83.50
		71.25
		
## Built With

* [munkres](https://github.com/clyphub/munkres) - an implementation of the Hungarian algorithm

## Assumptions

- Input file MUST be ASCII encoded
- For each input line, the first `;` encountered will be used as the delimiter, if there are more `;`s in the line, they become part of the product name







