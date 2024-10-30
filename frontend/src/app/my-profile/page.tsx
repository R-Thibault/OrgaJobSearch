"use client";

import axios from "axios";
import React, { useEffect, useState } from "react";

export default function MyProfilePage() {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [userName, setUserName] = useState<string>();
  const [userEmail, setUserEmail] = useState<string>();
  const [successMessage, setSuccessMessage] = useState<string>();
  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const response = await axios.get("http://localhost:8080/me", {
          withCredentials: true,
        });
        if (response.status === 200) {
          console.log(response.data);
          setUserName(response.data.userName);
          setUserEmail(response.data.userEmail);
        } else {
          setUserName("");
          setUserEmail("");
          setErrorMessage(
            "Une érreur est survenue, veuillez contactez un administrateur"
          );
        }
      } catch (error) {
        console.error("Error during profile setup:", error);
        setErrorMessage(
          "Une érreur est survenue, veuillez contactez un administrateur"
        );
      }
    };

    fetchProfile();
  }, []);
  const handleSubmit = async () => {
    const response = await axios.post(
      "http://localhost:8080/update-user",
      {
        UserName: userName,
        Email: userEmail,
      },
      { withCredentials: true }
    );
    if (response.status == 200) {
      setSuccessMessage("Mise a jour de votre profile effectué avec succés");
    }
  };
  if (errorMessage != null) {
    return <div className="text-red-500 text-sm mb-4">{errorMessage}</div>;
  } else {
    return (
      <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center">
        <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
          <h1 className="text-2xl font-bold mb-4">My Profile</h1>
          {successMessage && (
            <span className="text-green-500 text-sm mb-4">
              {successMessage}
            </span>
          )}
          <div className="mb-4">
            <label className="block text-gray-700">Name</label>
            <input
              type="text"
              className="mt-1 p-2 w-full border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Your Name"
              value={userName}
              onChange={(e) => setUserName(e.target.value)}
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Email</label>
            <input
              type="email"
              className="mt-1 p-2 w-full border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={userEmail}
              readOnly
            />
          </div>
          <button
            className="w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
            type="button"
            onClick={handleSubmit}
          >
            Save
          </button>
        </div>
      </div>
    );
  }
}
