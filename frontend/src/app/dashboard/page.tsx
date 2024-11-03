"use client";
import { useEffect, useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";
import { FaExternalLinkAlt } from "react-icons/fa";

export default function Dashboard() {
  const [showAppModal, setShowAppModal] = useState<boolean>(false);
  const [appData, setAppData] = useState({
    url: "",
    title: "",
    description: "",
    location: "",
    company: "",
    applied: true,
  });
  interface Application {
    ID: string;
    Title: string;
    Url: string;
    Company: string;
    Location: string;
    Salary: number;
    Description: string;
    Applied: boolean;
    FollowUp: boolean;
    Response: boolean;
  }

  const [applications, setApplications] = useState<Application[]>([]);
  const router = useRouter();
  console.log("/dashboard");
  const handleLogout = async () => {
    try {
      const response = await axios.post(
        "http://localhost:8080/logout",
        {},
        { withCredentials: true }
      );

      if (response.data && response.data.message === "Logout successful") {
        router.push("/login");
      }
    } catch (error) {
      console.error("Error during Logout:", error);
      alert("Failed to Logout");
    }
  };

  const handleAppChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setAppData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleAppSubmit = async () => {
    // Here you can handle submitting the form data
    console.log("Application Data: ", appData);
    try {
      const response = await axios.post(
        "http://localhost:8080/application",
        {
          ...appData,
        },
        { withCredentials: true }
      );

      if (response.data === "message") {
        console.log(response);
      }
    } catch (error) {
      console.log(error);
    }
    setShowAppModal(false);
  };

  useEffect(() => {
    const fetchApplications = async () => {
      try {
        const response = await axios.get(
          "http://localhost:8080/get-applications-by-user",
          { withCredentials: true }
        );
        if (response.status === 200) {
          setApplications(response.data.message);
        }
      } catch (error) {
        console.log(error);
      }
    };
    fetchApplications();
  }, []);
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
          <div>
            {/* Button to trigger application modal */}
            <button
              className="bg-purple-500 text-white px-4 py-2 rounded-md hover:bg-purple-600 ml-4"
              onClick={() => setShowAppModal(true)}
            >
              Create Application
            </button>
          </div>

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
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {applications.length > 0 ? (
            applications.map((app) => (
              <div
                key={app.ID}
                className="p-4 bg-gray-100 border rounded-md shadow-lg transition-all transform hover:scale-105 hover:shadow-2xl relative"
              >
                {/* Visible status */}
                <div className="flex items-center justify-between">
                  <h3 className="text-lg font-semibold">{app.Title}</h3>
                  <a
                    href={app.Url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-blue-600 hover:text-blue-800"
                    aria-label="View Job Posting"
                  >
                    <FaExternalLinkAlt />
                  </a>
                </div>
                <div className="text-sm font-semibold mb-2">Status</div>
                <p className="text-sm mb-2">
                  Applied: {app.Applied ? "Yes" : "No"}
                </p>
                <p className="text-sm mb-2">
                  Response: {app.Response ? "Yes" : "No"}
                </p>
                <p className="text-sm mb-4">
                  Follow Up: {app.FollowUp ? "Yes" : "No"}
                </p>

                {/* Full hover overlay */}
                <div className="absolute inset-0 p-4 bg-white rounded-md opacity-0 hover:opacity-100 transition-opacity duration-300">
                  <div className="flex items-center justify-between">
                    <h3 className="text-lg font-semibold">{app.Title}</h3>
                    <a
                      href={app.Url}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-blue-600 hover:text-blue-800"
                      aria-label="View Job Posting"
                    >
                      <FaExternalLinkAlt />
                    </a>
                  </div>
                  <p className="text-gray-600 mt-2">Company: {app.Company}</p>
                  <p className="text-gray-600">Location: {app.Location}</p>
                  <p className="text-gray-600">Salary: ${app.Salary}</p>
                  <p className="text-gray-600 mt-2 line-clamp-2">
                    {app.Description}
                  </p>
                </div>
              </div>
            ))
          ) : (
            <p>No applications found.</p>
          )}
        </div>
      </div>

      {/* Modal for creating an application */}
      {showAppModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
          <div className="relative bg-white p-6 rounded-md shadow-md w-96">
            {/* Close button */}
            <button
              className="absolute top-4 right-4 text-gray-500 hover:text-gray-700 text-2xl font-bold"
              onClick={() => setShowAppModal(false)}
              aria-label="Close"
            >
              &times;
            </button>
            <h2 className="text-xl font-semibold mb-4">Create Application</h2>
            <form onSubmit={handleAppSubmit}>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">
                  URL<span className="text-red-600"> *</span>
                  <input
                    type="text"
                    name="url"
                    value={appData.url}
                    id="url"
                    placeholder=" "
                    onChange={handleAppChange}
                    className="w-full px-3 py-2 border rounded-md required:border-red-500 valid:border-green-500"
                    required
                  />
                </label>
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">
                  Title<span className="text-red-600"> *</span>
                </label>
                <input
                  type="text"
                  name="title"
                  value={appData.title}
                  onChange={handleAppChange}
                  className="w-full px-3 py-2 border rounded-md required:border-red-500 valid:border-green-500"
                  required
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">
                  Description
                </label>
                <textarea
                  name="description"
                  value={appData.description}
                  onChange={handleAppChange}
                  className="w-full px-3 py-2 border rounded-md "
                ></textarea>
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">
                  Location
                </label>
                <input
                  type="text"
                  name="location"
                  value={appData.location}
                  onChange={handleAppChange}
                  className="w-full px-3 py-2 border rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">
                  Company
                </label>
                <input
                  type="text"
                  name="company"
                  value={appData.company}
                  onChange={handleAppChange}
                  className="w-full px-3 py-2 border rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">
                  Applied
                </label>
                <div className="flex items-center">
                  <label className="mr-4">
                    <input
                      type="radio"
                      name="applied"
                      value="yes"
                      checked={appData.applied === true}
                      onChange={handleAppChange}
                      defaultValue={"yes"}
                    />
                    Yes
                  </label>
                  <label>
                    <input
                      type="radio"
                      name="applied"
                      value="no"
                      checked={appData.applied === false}
                      onChange={handleAppChange}
                    />
                    No
                  </label>
                </div>
              </div>
              <div className="flex justify-end mt-4">
                <button
                  className="bg-gray-300 text-black px-4 py-2 rounded-md mr-2 hover:bg-gray-400"
                  onClick={() => setShowAppModal(false)}
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
                >
                  Create Application
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
