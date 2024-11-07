"use client";
import React, { useState } from "react";

// Import des composants pour les deux types d'inscription
import SignUpWithOTP from "../../components/organisms/signUpWithOTP";

export default function SignUpPage() {
  const [email, setEmail] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  // Afficher le composant approprié en fonction de si l'OTP est requis ou non
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <SignUpWithOTP
        initialEmail={email}
        firstName={firstName}
        setFirstName={setFirstName}
        lastName={lastName}
        setLastName={setLastName}
        password={password}
        setPassword={setPassword}
        confirmPassword={confirmPassword}
        setConfirmPassword={setConfirmPassword}
      />
    </div>
  );
}
