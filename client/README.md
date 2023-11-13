# To Do

- Add CSS Styling to Components
- Add Flash Message for success message
    - Login
    - Sign Up
- Change Cookie JWT Token to sessionstorage JWT Token.

# Issues Faced

1. Cookie was not set in browser cookie but can be seen in login response network (under 'Cookies').
    - have to set `'credentials': "include"` in `fetch` request that is expecting a cookie in the response.
    - For axios, use `withCredentials: true`.
    - In Golang CORS enabling, ensure that this field is set to true:
        - `AllowCredentials: true`

# TypeScript

## `React.ReactNode`

## `React.ReactElement`

## `React.ReactFC<Props>`

## Generics in TypeScript React

- [Link](https://devtrium.com/posts/react-typescript-using-generics-in-react)
- Generics provide a way to tell functions, classes, or interfaces what type you want to use when you call it.

```js
// Generics
function identity<ArgType>(arg: ArgType): ArgType {
  return arg;
}

// How we would use it
const greeting = identity<string>('Hello World!');

```