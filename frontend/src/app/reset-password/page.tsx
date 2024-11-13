"use client";
import React, { useEffect, useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";

export default function ResetPassword() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [isTokenValid, setIsTokenValid] = useState<boolean | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [message, setMessage] = useState("");
  const [checkPswd, setCheckPswnd] = useState(
    "Au moins 8 caractères, avec une majuscule, une minuscule, un chiffre et un caractère spécial (@ $ ! % * ? &)."
  );
  const passwordRegex =
    /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const extractedToken = params.get("token");
    if (extractedToken) {
      setToken(extractedToken);
      sendTokenToBackend(extractedToken);
    } else {
      setIsTokenValid(false);
    }
  }, []);
  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setPassword(value);
    if (passwordRegex.test(value)) {
      setCheckPswnd("");
    } else {
      setCheckPswnd(
        "Le mot de passe doit contenir au moins 8 caractères, avec une majuscule, une minuscule, un chiffre et un caractère spécial (@ $ ! % * ? &)."
      );
    }
  };
  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  };

  // Function to send the extracted token to the backend
  const sendTokenToBackend = async (token: string) => {
    try {
      const response = await axios.post(
        "http://localhost:8080/verify-reset-password-link",
        {
          token,
        }
      );
      if (response.data) {
        console.log(response.data);
        setEmail(response.data.message);
        setIsTokenValid(true);
      } else {
        setErrorMessage("Invalid or expired token. Please contact support.");
        setIsTokenValid(false);
      }
    } catch (error) {
      console.log("Error verifying token:", error);
      setErrorMessage("Erreur lors de la vérification du token.");
      setIsTokenValid(false);
    }
  };
  const handleResetPassword = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setMessage("");
    setErrorMessage("");

    try {
      const response = await axios.post(
        "http://localhost:8080/reset-password",
        {
          email,
          password,
          confirmPassword,
          tokenString: token,
        }
      );
      console.log("TEST");
      if (response.status === 200) {
        console.log("200", response);
        setMessage("The reset password email has been sent successfully.");
      } else {
        console.log("other", response);
        setErrorMessage("An error occurred. Please contact the administrator.");
      }
    } catch (err) {
      console.error("Error sending reset password email:", err);
      setErrorMessage("An error occurred. Please contact the administrator.");
    }
  };
  if (isTokenValid === null) {
    return <p className="text-center">Vérification du token en cours...</p>;
  } else if (isTokenValid === false) {
    return <p className="text-red-500 text-center">{errorMessage}</p>;
  } else {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <div className="bg-white p-8 rounded shadow-md w-full max-w-md">
          <h2 className="text-2xl font-bold mb-6 text-center">
            Reset Password
          </h2>
          <form className="space-y-4" onSubmit={handleResetPassword}>
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                Email
              </label>
              <input
                type="email"
                id="email"
                className="w-full px-3 py-2 mt-1 border rounded-md shadow-sm bg-gray-100 sm:text-sm focus:outline-none focus:border-gray-300 hover:border-gray-300 cursor-not-allowed"
                value={email}
                onChange={handleEmailChange}
                required
                readOnly
              />
            </div>
            <div>
              <label
                htmlFor="password"
                className="block text-sm font-medium text-gray-700"
              >
                Password
              </label>
              <input
                type="password"
                id="password"
                className="w-full px-3 py-2 mt-1 border rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                value={password}
                onChange={handlePasswordChange}
                required
              />
            </div>

            {checkPswd && (
              <p className="text-red-500 text-sm mb-4">{checkPswd}</p>
            )}
            <div>
              <label
                htmlFor="confirmPassword"
                className="block text-sm font-medium text-gray-700"
              >
                Confirm Password
              </label>
              <input
                type="password"
                id="confirmPassword"
                className="w-full px-3 py-2 mt-1 border rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                required
              />
            </div>
            {message && (
              <p className="text-green-500 text-sm mb-4">{message}</p>
            )}
            {errorMessage && (
              <p className="text-red-500 text-sm mb-4">{errorMessage}</p>
            )}
            <button
              type="submit"
              className="w-full px-4 py-2 text-white bg-indigo-600 rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Submit
            </button>
          </form>
        </div>
      </div>
    );
  }
}
