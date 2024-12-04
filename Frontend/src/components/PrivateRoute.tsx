import { Navigate } from "react-router-dom";

import { useCookies } from "react-cookie";
import React from "react";

interface ProtectedRouteProps {
    children: JSX.Element;
}
  

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {

    //eslint-disable-next-line
	const [cookie] = useCookies();

    if (cookie["token"]) {
        return children;
    }

    return <Navigate to="/login" />
};