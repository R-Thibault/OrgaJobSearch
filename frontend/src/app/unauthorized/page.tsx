// /src/app/unauthorized.tsx

import Link from "next/link";

export default function Unauthorized() {
  return (
    <div className="flex items-center justify-center h-screen bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md max-w-md text-center">
        <h1 className="text-3xl font-bold text-red-600 mb-4">Access Denied</h1>
        <p className="text-gray-700 mb-6">
          You don’t have permission to access this page.
        </p>
        <div className="flex justify-center gap-4">
          <Link href="/">
            <a className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors">
              Go to Home
            </a>
          </Link>
          <Link href="/login">
            <a className="bg-gray-300 text-gray-800 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors">
              Login
            </a>
          </Link>
        </div>
      </div>
    </div>
  );
}
