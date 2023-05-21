# To Do

- Add CSS Styling to Components
- Add Flash Message for success message
    - Login
    - Sign Up

# Issues Faced

1. Cookie was not set in browser cookie but can be seen in login response network (under 'Cookies').
    - have to set `'credentials': "include"` in `fetch` request that is expecting a cookie in the response.
    - For axios, use `withCredentials: true`.
    - In Golang CORS enabling, ensure that this field is set to true:
        - `AllowCredentials: true`