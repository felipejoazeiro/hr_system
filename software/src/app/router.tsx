import {createBrowserRouter, RouterProvider} from "react-router-dom"
import LoginPage from "../features/login/page/LoginPage";


const router = createBrowserRouter([
    {
        path: "/", element: <LoginPage />,
    }
]);

export default function AppRouter(){ 
    return <RouterProvider router={router} />;
}