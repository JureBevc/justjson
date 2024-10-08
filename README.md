# JustJSON
The best frontend framework based on the best file format. 

Why bother with messy HTML, CSS, or JavaScript when you can build entire websites using only **JSON**? Say goodbye to the tyranny of traditional web development and say hello to the simplicity of objects and arrays!


## Description
Are you tired of learning multiple languages like HTML, CSS, and JavaScript? Do you feel overwhelmed by the modern frontend ecosystem, with its endless layers of abstraction and package managers? I know you do. That's why I built **JustJSON**—the framework that compiles pure JSON into an entire website.

Now you can finally build your dream website with *just* JSON! Who needs JSX, templating languages, or even a browser? JustJSON is so revolutionary, you don’t even need to know what a browser is! If you’re smart enough to write JSON, you’re smart enough to write **everything**.


It is so good, yet so simple, you will wonder how come no one thought of this before. You're welcome.

## Features
- 🚀 **Zero setup**: No npm, no webpack, no bulky tools—just your favorite text editor, a JSON file and JustJSON.
- 💪 **Frontend in JSON**: One file format to rule them all!
- 🔥 **Performance**: Caring about performance in web development is an anti-pattern, just get the feature out ASAP with JustJSON!
- 🛠️ **No learning curve**: If you know how to make a JSON file, congratulations—you already mastered JustJSON!
- 🥇 **Best practices**: Yes, you will practice JSON all the time.
- 🌍 **Widespread adoption**: 100% of developers in my home office use and enjoy JustJSON.

## Usage
To build your website with JustJSON, simply run the following command.
```
./JustJSON ./website.json
```

This will generate HTML, CSS and Javascript files in the `./build` directory.


## Examples

You can check out the [example](https://github.com/JureBevc/justjson/tree/main/example) folder to see the full demo, but here are some snippets to give you an idea what's it like programming with JustJSON.

### HTML

```json
{
    "definitions": [
        {
            "type": "html",
            "name": "index.html",
            "elements": [
                {
                    "tag": "head",
                    "elements": [
                        {
                            "tag": "title",
                            "elements": [
                                "My JSON Website"
                            ]
                        }
                    ]
                },
                {
                    "tag": "body",
                    "elements": [
                        {
                            "tag": "h1",
                            "elements": [
                                "I built this with JSON!"
                            ]
                        }
                    ]
                }
            ]
        }
    ]
}
```

### CSS

```json
{
    "type": "css",
    "name": "style.css",
    "content": {
        "body": {
            "background-color": "#fefefe",
            "font-family": "'Comic Sans MS', cursive, sans-serif"
        },
        "h1": {
            "color": "rebeccapurple"
        }
    }
}
```

### JavaScript

```json
{
    "type": "javascript",
    "name": "script.js",
    "commands": [
        {
            "type": "let",
            "name": "a",
            "value": "0"
        },
        {
            "type": "let",
            "name": "b",
            "value": "1"
        },
        {
            "type": "let",
            "name": "result",
            "value": "0"
        },
        {
            "type": "function",
            "name": "fibonacci",
            "parameters": [],
            "commands": [
                {
                    "type": "set",
                    "variable": "result",
                    "value": {
                        "type": "operator",
                        "operator": "+",
                        "left": "a",
                        "right": "b"
                    }
                },
                {
                    "type": "set",
                    "variable": "a",
                    "value": "b"
                },
                {
                    "type": "set",
                    "variable": "b",
                    "value": "result"
                },
                {
                    "type": "set",
                    "variable": "document.getElementById('fibonacci-display').innerText",
                    "value": "result"
                }
            ]
        },
        {
            "type": "call",
            "function": "document.getElementById('fibonacci-btn').addEventListener",
            "parameters": [
                "'click'",
                "fibonacci"
            ]
        }
    ]
}
```