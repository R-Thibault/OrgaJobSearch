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
    <div className="flex flex-col items-center justify-center gap-4">
      {applications.length > 0 ? (
        applications.map((app) => (
          <div key={app.ID} className="w-full max-w-2xl">
            <ApplicationCard {...app} />
          </div>
        ))
      ) : (
        <p>No applications found.</p>
      )}
    </div>
  );
}
