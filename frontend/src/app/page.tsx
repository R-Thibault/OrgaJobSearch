import Navbar from "../components/organisms/navbar";
export default function Home() {
  return (
    <div className="min-h-screen bg-gray-100 flex flex-col justify-between">
      <Navbar />
      <main className="container mx-auto p-8">
        <h1 className="text-3xl font-bold text-gray-800 mb-4">
          Welcome to MyApp
        </h1>
        <p className="text-gray-700 mb-4">
          This is a simple example of a Next.js page with a Tailwind CSS navbar.
          The page is styled using Tailwind CSS utility classes, providing a
          modern, responsive layout with minimal effort.
        </p>
        <p className="text-gray-700">
          Explore the navigation links above to learn more about what this site
          has to offer.
        </p>
      </main>
      <footer className="bg-blue-600 p-4 mt-8">
        <div className="container mx-auto text-center text-white">
          <p>&copy; 2024 MyApp. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
}
