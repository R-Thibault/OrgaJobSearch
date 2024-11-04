" use client";

import { ApplicationType } from "@/types/applicationType";
import { FaExternalLinkAlt } from "react-icons/fa";

export default function ApplicationCard(application: ApplicationType) {
  return (
    <div className="p-4 bg-gray-100 border rounded-md shadow-lg transition-all transform hover:scale-105 hover:shadow-2xl relative">
      {/* Visible status */}
      <div className="flex items-center ">
        <h3 className="text-lg font-semibold">{application.Title}</h3>
        <a
          href={application.Url}
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
        Applied: {application.Applied ? "Yes" : "No"}
      </p>
      <p className="text-sm mb-2">
        Response: {application.Response ? "Yes" : "No"}
      </p>
      <p className="text-sm mb-4">
        Follow Up: {application.FollowUp ? "Yes" : "No"}
      </p>

      {/* Full hover overlay */}
      <div className="absolute inset-0 p-4 bg-white rounded-md opacity-0 hover:opacity-100 transition-opacity duration-300">
        <div className="flex items-center justify-between">
          <h3 className="text-lg font-semibold">{application.Title}</h3>
          <a
            href={application.Url}
            target="_blank"
            rel="noopener noreferrer"
            className="text-blue-600 hover:text-blue-800"
            aria-label="View Job Posting"
          >
            <FaExternalLinkAlt />
          </a>
        </div>
        <p className="text-gray-600 mt-2">Company: {application.Company}</p>
        <p className="text-gray-600">Location: {application.Location}</p>
        <p className="text-gray-600">Salary: ${application.Salary}</p>
        <p className="text-gray-600 mt-2 line-clamp-2">
          {application.Description}
        </p>
      </div>
    </div>
  );
}
