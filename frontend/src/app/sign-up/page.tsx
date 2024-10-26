"use client";
import React, { useEffect, useState } from "react";
import axios from "axios";

// Import des composants pour les deux types d'inscription
import SignUpWithOTP from "../../components/organisms/signUpWithOTP";
import SignUpWithoutOTP from "../../components/organisms/signUpWithoutOTP";

export default function SignUpPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [isTokenValid, setIsTokenValid] = useState<boolean | null>(null);
  const [isOtpRequired, setIsOtpRequired] = useState<boolean | null>(null);
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
      console.log(response.data);
      if (response.data) {
        setIsTokenValid(true);

        if (response.data.tokenType === "GlobalInvitation") {
          setIsOtpRequired(true);
        } else if (response.data.tokenType === "PersonalInvitation") {
          setEmail(response.data.email);
          setIsOtpRequired(false);
        } else {
          setIsTokenValid(false);
          setErrorMessage(
            "Méthode d'inscription non valide, veuillez contacter un administrateur."
          );
        }
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

  if (isTokenValid === null) {
    return <p className="text-center">Vérification du token en cours...</p>;
  } else if (isTokenValid === false) {
    return <p className="text-red-500 text-center">{errorMessage}</p>;
  } else {
    // Afficher le composant approprié en fonction de si l'OTP est requis ou non
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        {isOtpRequired
          ? token && (
              <SignUpWithOTP
                initialEmail={email}
                password={password}
                setPassword={setPassword}
                confirmPassword={confirmPassword}
                setConfirmPassword={setConfirmPassword}
                token={token}
              />
            )
          : token && (
              <SignUpWithoutOTP
                email={email}
                password={password}
                setPassword={setPassword}
                confirmPassword={confirmPassword}
                setConfirmPassword={setConfirmPassword}
                token={token}
              />
            )}
      </div>
    );
  }
}
