"use client";
import React, { useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";

interface SignUpWithoutOTPProps {
  email: string;
  password: string;
  setPassword: (password: string) => void;
  confirmPassword: string;
  setConfirmPassword: (confirmPassword: string) => void;
  token: string;
}

function SignUpWithoutOTP({
  email,
  password,
  setPassword,
  confirmPassword,
  setConfirmPassword,
  token,
}: SignUpWithoutOTPProps) {
  const [errorMessage, setErrorMessage] = useState("");
  const [checkPswd, setCheckPswnd] = useState(
    "Au moins 8 caractères, avec une majuscule, une minuscule, un chiffre et un caractère spécial (@ $ ! % * ? &)."
  );
  const router = useRouter();
  const passwordRegex =
    /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
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

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      if (password === confirmPassword) {
        const response = await axios.post("http://localhost:8080/sign-up", {
          email: email,
          password: password,
          tokenString: token,
        });
        if (response.data) {
          console.log(response.data);
          setErrorMessage("");
          router.push("/login");
        }
      } else {
        setErrorMessage(
          "Le mot de passe et la confirmation ne correspondent pas."
        );
      }
    } catch (error) {
      console.error("Error during sign up:", error);
      setErrorMessage("Erreur lors de l'inscription.");
    }
  };

  return (
    <div className="w-full max-w-md p-8 space-y-6 bg-white rounded shadow-md">
      <h2 className="text-2xl font-bold text-center">Sign Up Without OTP</h2>
      {errorMessage && (
        <p className="text-red-500 text-sm mb-4">{errorMessage}</p>
      )}
      <form className="space-y-4" onSubmit={handleSubmit}>
        <input type="hidden" name="token" value={token || ""} />
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
            className="w-full px-3 py-2 mt-1 border rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            value={email}
            readOnly
            required
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
        {checkPswd && <p className="text-red-500 text-sm mb-4">{checkPswd}</p>}
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
        <button
          type="submit"
          className="w-full px-4 py-2 text-white bg-indigo-600 rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Sign Up
        </button>
      </form>
    </div>
  );
}

export default SignUpWithoutOTP;
