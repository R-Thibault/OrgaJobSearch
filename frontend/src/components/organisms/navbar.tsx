"use client"; // Ensures the Navbar can handle client-side actions if needed (though it's kept simple here)

import React from "react";
import { useUser } from "@/app/context/UserContext";
import DashboardNavbar from "../molecules/dashboardNavbar";
import BackOfficeNavbar from "../molecules/backOfficeNavbar";
import HomeNavbar from "../molecules/homeNavbar";

export default function Navbar() {
  const { roles, loading } = useUser();
  if (loading) {
    return <div>Loading...</div>;
  }
  if (roles?.includes("JobSeeker")) {
    return <DashboardNavbar />;
  } else if (
    roles?.includes("CareerCoach") ||
    roles?.includes("CareerSupportManager")
  ) {
    return <BackOfficeNavbar />;
  } else {
    return <HomeNavbar />;
  }
}
