---
description: >-
  The purpose of this area of the wiki is to document works in progress, or to
  let others know the status of your improvements.  Best practices and project
  guidelines are defined here.
---

# Developer Area

## General Coding Guidelines

We don't have strict coding standards, but these few guild lines will help. First a few of the more strict guidelines.

* Indentation is done in tabs not spaces. You can always change your editors tab stop to make it easier for you.
* `using namespace std;` is forbidden at file scope and strongly discouraged to be used at all. The std namespace is large and can cause problems blindingly including everything.

And for the less strict guidelines.

* You should break up your if they are getting too long, we don't have strict limits of 80 characters per line like some projects, but you need to remember not everyone has the same screen resolution as you.
* If you are editing a function or block of code you should try to follow the style already present, unless the readability of the function is sufficiently poor that you feel it would be best to change it.

Most importantly, make the code readable.

Further reading: [Linux Kernel Coding Style](https://www.kernel.org/doc/Documentation/CodingStyle) is mostly logical and not all that ridiculous, we don't adhere to all of the rules that they do, but it is a fairly good base.

