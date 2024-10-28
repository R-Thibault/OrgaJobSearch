"use client";
import { useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";

export default function Dashboard() {
  const [showModal, setShowModal] = useState<boolean>(false);
  const router = useRouter();
  const handleLogout = async () => {
    try {
      const response = await axios.post(
        "http://localhost:8080/logout",
        {},
        { withCredentials: true }
      );
      console.log(response);
      if (response.data && response.data.message === "Logout successful") {
        router.push("/login");
      }
    } catch (error) {
      console.error("Error during Logout:", error);
      alert("Failed to Logout");
    }
  };

  return (
    <div className="flex h-screen">
      {/* Sidebar Navigation */}
      <nav className="w-64 bg-gray-800 text-white p-4">
        <h2 className="text-xl font-semibold mb-6">Dashboard Navigation</h2>
        <ul>
          <li className="mb-4">Dashboard Home</li>
          <li className="mb-4">User Management</li>
          <li className="mb-4">Settings</li>
        </ul>
      </nav>

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
            onClick={handleLogout}
          >
            Logout
          </button>
        </header>

        {/* Main content here */}
        <div className="text-xl">Welcome to the Dashboard!</div>
      </div>

      {/* Modal for sending invitation */}
      {showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
          <div className="bg-white p-6 rounded-md shadow-md w-96">
            <h2 className="text-xl font-semibold mb-4">Add application</h2>
            <span>Here add application form, work in progress</span>
            <div className="flex justify-end">
              <button
                className="bg-gray-300 text-black px-4 py-2 rounded-md mr-2 hover:bg-gray-400"
                onClick={() => setShowModal(false)}
              >
                Cancel
              </button>
              <button className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600">
                Send Invitation
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
