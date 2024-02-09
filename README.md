# uwuttp

## Description

uwuttp is a fun web server that UwU-ifies any given piece of text. Want to add a touch of cuteness and
playfulness to your sentences? Look no further! With uwuttp, you can transform your text into a delightful
UwU language, guaranteed to bring smiles and joy to anyone who reads it.

Whether you're a fan of OwO, UwU, or any other adorable expressions, uwuttp has got you covered. Simply provide
your text, and let uwuttp work its magic, turning it into a UwU-tastic masterpiece. With customizable options, uwuttp will make your text stand out and bring a sense of whimsy to your digital
interactions.

## Features

- UwU-ify any given piece of text.
- Lowercase the text for a softer and more adorable touch.
- Replace letters with their UwU counterparts.
- Add cute Unicode faces to express emotions.
- Introduce word stutters for extra cuteness.
- Customizable configuration options to tailor the UwU-ification process to your preferences.

## Usage

To use uwuttp, you need to send a GET request to the `/uwu` endpoint with the desired text and
configuration options as query parameters. Here's an example of how it can be used:

```http
GET /uwu?text=Hello+world&lower=true&replace=true&facesChance=50&stutterChance=25&stutterAmount=2
```

Here's an overview of the query parameters:

- `text`: The text you want to UwU-ify.
- `lower`: Set this to `true` if you want the text to be lowercased before UwU-ification. (default: `true`)
- `replace`: Set this to `true` if you want letters to be replaced with their UwU counterparts. (default: `true`)
- `facesChance`: The chance (in percentage) of adding cute Unicode faces to the text. Set a value between 0 and 100. (default: `10`)
- `stutterChance`: The chance (in percentage) of a word being stuttered. Set a value between 0 and 100. (default: `30`)
- `stutterAmount`: The amount of times a word should be stuttered. Set a positive integer value. (default: `1`)

The server will respond with a JSON object containing the UwU-ified text.

## Installation

To install uwuttp, simply clone this repository to your local machine and build the project using your
preferred Go build tool. No additional dependencies are required.

```bash
git clone https://github.com/interrrp/uwuttp.git
cd uwuttp
go build
```

Alternatively, you could use Docker:

```bash
docker compose up --d --build
```

You can now access the server at <http://localhost:8080/> (see
[the Compose file](docker-compose.yml>)).
