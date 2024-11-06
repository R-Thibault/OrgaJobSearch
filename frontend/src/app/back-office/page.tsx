"use client";
import { useState } from "react";
import axios from "axios";

export default function BackOffice() {
  const [showModal, setShowModal] = useState<boolean>(false);
  const [showUrlModal, setShowUrlModal] = useState<boolean>(false);
  const [email, setEmail] = useState<string>("");
  const [generatedUrl, setGeneratedUrl] = useState<string>("");

  const handleSendInvitation = async () => {
    try {
      await axios.post(
        "http://localhost:8080/send-user-invitation",
        {
          email,
        },
        { withCredentials: true }
      );
      alert("Invitation sent successfully!");
      setShowModal(false);
      setEmail("");
    } catch (error) {
      console.error("Error sending invitation:", error);
      alert("Failed to send the invitation");
    }
  };

  const handleGenerateUrl = async () => {
    try {
      const response = await axios.post(
        "http://localhost:8080/generate-url",
        {
          invitationType: "GlobalInvitation",
        },
        { withCredentials: true }
      );

      setGeneratedUrl(response.data.url);
      setShowUrlModal(true);
    } catch (error) {
      console.log("Error generating URL:", error);
      alert("Failed to generate the URL");
    }
  };

  const handleCopyUrl = () => {
    navigator.clipboard.writeText(generatedUrl).then(
      () => {
        alert("URL copied to clipboard!");
      },
      (error) => {
        console.error("Error copying URL:", error);
        alert("Failed to copy the URL");
      }
    );
  };

  return (
    <div className="flex h-screen">
      {/* Main Content Area */}
      <div className="flex-1 p-6">
        <header className="flex items-center justify-between mb-8">
          {/* Button to trigger invitation modal */}
          <button
            className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
            onClick={() => setShowModal(true)}
          >
            Invite User
          </button>

          {/* Button to trigger URL modal */}
          <button
            className="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-600 ml-4"
            onClick={handleGenerateUrl}
          >
            Generate URL
          </button>
        </header>

        {/* Main content here */}
        <div className="text-xl">Welcome to the Dashboard!</div>
      </div>

      {/* Modal for sending invitation */}
      {showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
          <div className="bg-white p-6 rounded-md shadow-md w-96">
            <h2 className="text-xl font-semibold mb-4">Send Invitation</h2>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="Enter email address"
              className="border border-gray-300 p-2 rounded-md w-full mb-4"
            />
            <div className="flex justify-end">
              <button
                className="bg-gray-300 text-black px-4 py-2 rounded-md mr-2 hover:bg-gray-400"
                onClick={() => setShowModal(false)}
              >
                Cancel
              </button>
              <button
                className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
                onClick={handleSendInvitation}
              >
                Send Invitation
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Modal for displaying generated URL */}
      {showUrlModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
          <div className="bg-white p-6 rounded-md shadow-md w-96">
            <h2 className="text-xl font-semibold mb-4">Generated URL</h2>
            <input
              type="text"
              value={generatedUrl}
              readOnly
              className="border border-gray-300 p-2 rounded-md w-full mb-4"
            />
            <div className="flex justify-end">
              <button
                className="bg-gray-300 text-black px-4 py-2 rounded-md mr-2 hover:bg-gray-400"
                onClick={() => setShowUrlModal(false)}
              >
                Close
              </button>
              <button
                className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
                onClick={handleCopyUrl}
              >
                Copy URL
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
