" use client";

import { ApplicationType } from "@/types/applicationType";
import { FaExternalLinkAlt } from "react-icons/fa";

export default function ApplicationCard(application: ApplicationType) {
  return (
    <div className="w-full p-4 bg-white shadow-md rounded-lg p-4 border border-gray-200 transition-all transform hover:scale-105 hover:shadow-2xl relative">
      {/* Title and Link */}
      <div className="flex items-center p-4">
        <a
          href={application.Url}
          target="_blank"
          rel="noopener noreferrer"
          className="text-blue-600 hover:text-blue-800 mr-2"
          aria-label="View Job Posting"
        >
          <FaExternalLinkAlt />
        </a>
        <h3 className="text-lg font-semibold">{application.Title}</h3>
      </div>
      {/* Company and Location */}
      <div className="flex justify-between items-center px-4 text-gray-600 text-sm font-medium mb-2">
        <p>{application.Company}</p>
        <p>{application.Location}</p>
      </div>
      {/* Status */}
      <div className="flex justify-between px-4 pb-4 text-gray-600 text-sm font-semibold">
        <p>Applied: {application.Applied ? "Yes" : "No"}</p>
        <p>Response: {application.Response ? "Yes" : "No"}</p>
        <p>Follow Up: {application.FollowUp ? "Yes" : "No"}</p>
      </div>
      {/* Hover Overlay with Expanded Details */}
      <div className="absolute inset-0 p-4 bg-white rounded-md opacity-0 hover:opacity-100 transition-opacity duration-300 flex flex-col justify-between">
        {/* Title and Link */}
        <div className="flex items-center mb-2">
          <a
            href={application.Url}
            target="_blank"
            rel="noopener noreferrer"
            className="text-blue-600 hover:text-blue-800 mr-2"
            aria-label="View Job Posting"
          >
            <FaExternalLinkAlt />
          </a>
          <h3 className="text-lg font-semibold">{application.Title}</h3>
        </div>

        {/* Company, Location, and Description */}
        <div className="flex flex-col gap-2">
          <div className="flex justify-between text-gray-600 text-sm font-medium">
            <p>{application.Company}</p>
            <p>{application.Location}</p>
          </div>
          <p className="text-gray-600 text-sm mt-2 line-clamp-3">
            {application.Description}
          </p>
        </div>

        {/* Status (Aligned at the bottom) */}
        <div className="flex justify-between text-gray-600 text-sm font-semibold mt-4">
          <p>Applied: {application.Applied ? "Yes" : "No"}</p>
          <p>Response: {application.Response ? "Yes" : "No"}</p>
          <p>Follow Up: {application.FollowUp ? "Yes" : "No"}</p>
        </div>
      </div>
    </div>
  );
}
