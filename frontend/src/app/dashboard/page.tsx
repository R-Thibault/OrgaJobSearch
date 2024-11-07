"use client";
import { useEffect, useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";
import ApplicationDisplay from "@/components/organisms/applicationsDisplay";
import { ApplicationType } from "@/types/applicationType";
import ApplicationFormModal from "@/components/organisms/applicationFormModal";

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

  const [applications, setApplications] = useState<ApplicationType[]>([]);
  const router = useRouter();

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

  const handleAppSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        "http://localhost:8080/create-application",
        {
          ...appData,
        },
        { withCredentials: true }
      );

      if (response.status === 200) {
        fetchApplications();
      }
    } catch (error) {
      console.log(error);
    }
    setShowAppModal(false);
  };
  const fetchApplications = async () => {
    try {
      const response = await axios.post(
        "http://localhost:8080/get-applications-by-user",
        {
          limit: 10,
          offset: 0,
          order: "asc",
          where: "",
        },
        { withCredentials: true }
      );
      if (response.status === 200) {
        setApplications(response.data.message);
      }
    } catch (error) {
      console.log(error);
    }
  };
  useEffect(() => {
    fetchApplications();
  }, []);
  return (
    <div className="flex min-h-screen">
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
        <ApplicationDisplay applications={applications} />
      </div>

      {/* Modal for creating an application */}
      <ApplicationFormModal
        showModal={showAppModal}
        onClose={() => setShowAppModal(false)}
        onSubmit={handleAppSubmit}
        appData={appData}
        onChange={handleAppChange}
      />
    </div>
  );
}
