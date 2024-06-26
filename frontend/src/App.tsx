import { FC } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import TasksPage from "./pages/tasks";
import AgentsPage from "./pages/agents";
import TimeoutsPage from "./pages/timeouts";
import AuthPage from "./pages/auth";
import Navbar, { ILink } from "./components/navbar";
import Layout from "./components/layout";
import Sidebar from "./components/sidebar";
import { QueryClient, QueryClientProvider } from "react-query";

const links: ILink[] = [
  {
    to: "/",
    label: "tasks",
  },
  {
    to: "/agents",
    label: "agents",
  },
  {
    to: "/timeouts",
    label: "timeouts",
  },
];

const App: FC = () => {
  const queryClient = new QueryClient();
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Layout navbar={<Navbar links={links} />} sidebar={<Sidebar />}>
          <Routes>
            <Route path="/" element={<TasksPage />} />
            <Route path="/agents" element={<AgentsPage />} />
            <Route path="/timeouts" element={<TimeoutsPage />} />
            <Route path="/auth" element={<AuthPage />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </QueryClientProvider>
  );
};

export default App;
