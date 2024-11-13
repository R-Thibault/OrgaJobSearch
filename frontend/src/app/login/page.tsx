"use client";

import axios from "axios";
import { useRouter } from "next/navigation";
import React, { useState } from "react";

export default function SignIn() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [resetEmail, setResetEmail] = useState("");
  const [resetMessage, setResetMessage] = useState("");
  const [resetError, setResetError] = useState("");

  const router = useRouter();
  const handleSignIn = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        "http://localhost:8080/login",
        {
          email: email,
          password: password,
        },
        { withCredentials: true }
      );
      // Check if the response was successful (status 200)
      if (response.status === 200) {
        const responseMe = await axios.get("http://localhost:8080/me", {
          withCredentials: true,
        });
        if (responseMe.status === 200) {
          router.push("/dashboard");
          router.refresh();
        } else {
          setErrorMessage("An unexpected error occurred. Please try again.");
        }
      }
    } catch (error) {
      console.error("Error during sign in:", error);
      setErrorMessage("Email ou mot de passe incorrecte.");
    }
  };

  const handleResetPassword = async (e: React.FormEvent) => {
    e.preventDefault();
    setResetMessage("");
    setResetError("");

    try {
      const response = await axios.post(
        "http://localhost:8080/send-reset-password-link",
        {
          email: resetEmail,
        }
      );
      if (response.status === 200) {
        setResetMessage("The reset password email has been sent successfully.");
      } else {
        setResetError("An error occurred. Please contact the administrator.");
      }
    } catch (err) {
      console.error("Error sending reset password email:", err);
      setResetError("An error occurred. Please contact the administrator.");
    }
  };
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 rounded shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold mb-6 text-center">Sign In</h2>
        <form onSubmit={handleSignIn}>
          <div className="mb-4">
            <label
              className="block text-gray-700 text-sm font-bold mb-2"
              htmlFor="email"
            >
              Email
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="email"
              type="email"
              value={email}
              placeholder="Email"
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="mb-6">
            <label
              className="block text-gray-700 text-sm font-bold mb-2"
              htmlFor="password"
            >
              Password
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
              id="password"
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          {errorMessage && (
            <p className="text-red-500 text-sm mb-4">{errorMessage}</p>
          )}
          <div className="flex items-center justify-between">
            <button
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
            >
              Sign In
            </button>
            <button
              type="button"
              onClick={() => setIsModalOpen(true)}
              className="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
            >
              Forgot Password?
            </button>
          </div>
        </form>
      </div>
      {/* Modal for Reset Password */}
      {isModalOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
          <div className="bg-white p-6 rounded shadow-md w-full max-w-sm">
            <h3 className="text-xl font-bold mb-4 text-center">
              Reset Password
            </h3>
            <form onSubmit={handleResetPassword}>
              <label
                className="block text-gray-700 text-sm font-bold mb-2"
                htmlFor="resetEmail"
              >
                Email
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="resetEmail"
                type="email"
                value={resetEmail}
                placeholder="Enter your email"
                onChange={(e) => setResetEmail(e.target.value)}
                required
              />
              {resetMessage && (
                <p className="text-green-500 text-sm mt-4">{resetMessage}</p>
              )}
              {resetError && (
                <p className="text-red-500 text-sm mt-4">{resetError}</p>
              )}
              <div className="flex items-center justify-between mt-4">
                <button
                  type="button"
                  onClick={() => setIsModalOpen(false)}
                  className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                >
                  Return
                </button>
                <button
                  className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                  type="submit"
                >
                  Send Reset Email
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
