"use client";

import axios from "axios";
import { useRouter } from "next/navigation";
import React, { useState } from "react";

export default function SignIn() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const router = useRouter();
  const handleSignIn = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log("LOGIN");
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
          const { userRole } = responseMe.data;
          console.log(responseMe);
          if (
            userRole.includes("CareerSupportManager") ||
            userRole.includes("CareerCoach")
          ) {
            router.push("/back-office");
          } else {
            // Default or catch-all route
            router.push("/dashboard");
          }
          router.refresh();
        }
        // router.push("/dashboard");
        // router.refresh();
      } else {
        setErrorMessage("An unexpected error occurred. Please try again.");
      }
    } catch (error) {
      console.error("Error during sign in:", error);
      setErrorMessage("Email ou mot de passe incorrecte.");
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
          </div>
        </form>
      </div>
    </div>
  );
}
