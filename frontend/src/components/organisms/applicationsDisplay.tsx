"use client";
import { ApplicationType } from "@/types/applicationType";

import ApplicationCard from "../molecules/applicationCard";

interface ApplicationDisplayProps {
  applications: ApplicationType[];
}

export default function ApplicationDisplay({
  applications,
}: ApplicationDisplayProps) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {applications.length > 0 ? (
        applications.map((app) => <ApplicationCard key={app.ID} {...app} />)
      ) : (
        <p>No applications found.</p>
      )}
    </div>
  );
}
