import { FC } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import TasksPage from "./pages/tasks";
import AgentsPage from "./pages/agents";
import TimeoutsPage from "./pages/timeouts";
import AuthPage from "./pages/auth";
import Navbar from "./components/navbar";
import Layout from "./components/layout";
import Sidebar from "./components/sidebar";

const App: FC = () => {
  return (
    <BrowserRouter>
      <Layout navbar={<Navbar />} sidebar={<Sidebar />}>
        <Routes>
          <Route path="/" element={<TasksPage />} />
          <Route path="/agents" element={<AgentsPage />} />
          <Route path="/timeouts" element={<TimeoutsPage />} />
          <Route path="/auth" element={<AuthPage />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  );
};

export default App;
