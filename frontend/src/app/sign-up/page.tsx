"use client";
import React, { useEffect, useState } from "react";
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
  const [isTokenValid, setIsTokenValid] = useState<boolean | null>(null);
  const [otp, setOtp] = useState("");
  const [isOtpModalOpen, setIsOtpModalOpen] = useState(false);
  const router = useRouter();
  const passwordRegex =
    /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;

  // New state to store the token
  const [token, setToken] = useState<string | null>(null);

  // Extract the token from the URL when the page loads
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

  // Function to send the extracted token to the backend
  const sendTokenToBackend = async (token: string) => {
    try {
      const response = await axios.post("http://localhost:8080/verify-token", {
        token,
      });

      if (response.data && response.data.valid) {
        console.log("Token is valid.");
        setIsTokenValid(true);
        setEmail(response.data.email);
      } else {
        setErrorMessage("Invalid or expired token. Please contact support.");
        setIsTokenValid(false);
      }
    } catch (error) {
      console.error("Error verifying token:", error);
      setErrorMessage("Erreur lors de la vérification du token.");
      setIsTokenValid(false);
    }
  };

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
          token: token,
        });
        if (response.data) {
          console.log(response.data);
          setErrorMessage("");
          // const responseOTP = await axios.post(
          //   "http://localhost:8080/generate-otp",
          //   {
          //     email: email,
          //   }
          // );
          // console.log(responseOTP);
          // // Open OTP modal after successful signup
          // setIsOtpModalOpen(true);
          // if (responseOTP.data) {
          //   setErrorMessage("");
          //   setErrorOTPMessage("");
          // } else {
          //   setErrorOTPMessage(
          //     "Erreur lors de l'envoi de l'email, veuillez contacter un administrateur"
          //   );
          // }
          router.push("/login");
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
      {isTokenValid === null ? (
        <p className="text-center">Vérification du token en cours...</p>
      ) : isTokenValid === false ? (
        <p className="text-red-500 text-center">{errorMessage}</p>
      ) : (
        <div className="w-full max-w-md p-8 space-y-6 bg-white rounded shadow-md">
          <h2 className="text-2xl font-bold text-center">Sign Up</h2>
          {errorMessage && (
            <p className="text-red-500 text-sm mb-4">{errorMessage}</p>
          )}
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
            <input type="hidden" name="token" value={token || ""} />
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
          {/* OTP Modal */}
          {isOtpModalOpen && (
            <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
              <div className="w-full max-w-md p-6 bg-white rounded-md shadow-lg">
                <h3 className="text-xl font-bold text-center mb-4">
                  Enter OTP
                </h3>
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
                    <p className="text-red-500 text-sm mb-4">
                      {errorOTPMessage}
                    </p>
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
      )}
    </div>
  );
}
