// src/context/AuthContext.js
import React, { useState, useEffect } from "react";
import { jwtDecode } from "jwt-decode";

const AuthContext = React.createContext();

export const AuthProvider = ({ children }) => {
    const [token, setToken] = useState(localStorage.getItem("token") || null);
    const [userId, setUserId] = useState(null);
    const [userRole, setUserRole] = useState(null); // Guardar el rol del usuario
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    useEffect(() => {
        if (token) {
            try {
                const decoded = jwtDecode(token);
                console.log("Decoded token:", decoded); // Verificar el contenido del token
                setUserId(decoded.user_id);
                setUserRole(decoded.role); // Asegurarse de que el rol esté en el token y guardarlo
                setIsAuthenticated(true);
            } catch (error) {
                console.error("Error decoding token:", error);
                setToken(null);
                setUserId(null);
                setUserRole(null);
                setIsAuthenticated(false);
            }
        }
    }, [token]);

    const login = (newToken) => {
        setToken(newToken);
        localStorage.setItem("token", newToken);
        try {
            const decoded = jwtDecode(newToken);
            setUserId(decoded.user_id);
            setUserRole(decoded.role); // Guardar el rol al iniciar sesión
            setIsAuthenticated(true);
        } catch (error) {
            console.error("Error decoding token on login:", error);
            setUserId(null);
            setUserRole(null);
            setIsAuthenticated(false);
        }
    };

    const logout = () => {
        setToken(null);
        setUserId(null);
        setUserRole(null);
        setIsAuthenticated(false);
        localStorage.removeItem("token");
    };

    return (
        <AuthContext.Provider value={{ token, userId, userRole, isAuthenticated, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export default AuthContext;