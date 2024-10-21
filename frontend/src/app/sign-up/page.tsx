"use client";
import React, { useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";

export default function SignUpPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [errorOTPMessage, setErrorOTPMessage] = useState("");
  const [checkPswd, setCheckPswnd] = useState(
    "Au moins 8 caractères, avec une majuscule, une minuscule, un chiffre et un caractère spécial (@ $ ! % * ? &)."
  );
  const [otp, setOtp] = useState("");
  const [isOtpModalOpen, setIsOtpModalOpen] = useState(false);
  const router = useRouter();
  const passwordRegex =
    /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;

  const handlePasswordChange = (e) => {
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
        });
        if (response.data) {
          console.log(response.data);
          setErrorMessage("");
          const responseOTP = await axios.post(
            "http://localhost:8080/generate-otp",
            {
              email: email,
            }
          );
          console.log(responseOTP);
          // Open OTP modal after successful signup
          setIsOtpModalOpen(true);
          if (responseOTP.data) {
            setErrorMessage("");
            setErrorOTPMessage("");
          } else {
            setErrorOTPMessage(
              "Erreur lors de l'envoi de l'email, veuillez contacter un administrateur"
            );
          }
        }
      } else {
        setErrorMessage(
          "Le mot de passe et la confirmation ne correspondent pas."
        );
      }
    } catch (error) {
      console.error("Error during sign up:", error);
      setErrorMessage("Email ou mot de passe incorrecte.");
    }
  };

  const handleOtpSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const otpResponse = await axios.post("http://localhost:8080/verify-otp", {
        email: email,
        otpCode: otp,
      });
      if (otpResponse.data) {
        setErrorOTPMessage("");
        // OTP is correct, proceed to login page
        router.push("/login");
      } else {
        console.log(otpResponse);
        setErrorOTPMessage("OTP incorrect. Veuillez réessayer.");
      }
    } catch (error) {
      console.error("Error during OTP verification:", error);
      setErrorOTPMessage("Erreur lors de la vérification de l'OTP.");
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-6 bg-white rounded shadow-md">
        <h2 className="text-2xl font-bold text-center">Sign Up</h2>
        <form className="space-y-4" onSubmit={handleSubmit}>
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
              onChange={(e) => setEmail(e.target.value)}
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
          {errorMessage && (
            <p className="text-red-500 text-sm mb-4">{errorMessage}</p>
          )}
          <button
            type="submit"
            className="w-full px-4 py-2 text-white bg-indigo-600 rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Sign Up
          </button>
        </form>
      </div>

      {/* OTP Modal */}
      {isOtpModalOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="w-full max-w-md p-6 bg-white rounded-md shadow-lg">
            <h3 className="text-xl font-bold text-center mb-4">Enter OTP</h3>
            <form onSubmit={handleOtpSubmit}>
              <div>
                <label
                  htmlFor="otp"
                  className="block text-sm font-medium text-gray-700"
                >
                  OTP
                </label>
                <input
                  type="text"
                  id="otp"
                  className="w-full px-3 py-2 mt-1 border rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  value={otp}
                  onChange={(e) => setOtp(e.target.value)}
                  required
                />
              </div>
              {errorOTPMessage && (
                <p className="text-red-500 text-sm mb-4">{errorOTPMessage}</p>
              )}
              <button
                type="submit"
                className="w-full px-4 py-2 mt-4 text-white bg-indigo-600 rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                Verify OTP
              </button>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
