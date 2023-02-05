import React, { Fragment } from "react";
import "./Auth.css";

function Auth() {
  return (
    <Fragment>
      <div className="auth-container">
        <h2 className="auth-header">Login</h2>
        <form className="auth-form">
          <div className="auth-form-group">
            <label htmlFor="username">Username:</label>
            <input
              type="text"
              placeholder="Username"
              name="username"
              autoFocus
              className="auth-input"
            />
          </div>
          <div className="auth-form-group">
            <label htmlFor="password">Password:</label>
            <input
              type="password"
              placeholder="Password"
              name="password"
              className="auth-input"
            />
          </div>
        </form>
      </div>
    </Fragment>
  );
}

export default Auth;
