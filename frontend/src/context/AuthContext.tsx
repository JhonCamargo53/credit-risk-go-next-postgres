'use client'
import { JWT_COOKIE_NAME } from '@/config/env.config';
import { LoginForm } from '@/types/auth';
import { User } from '@/types/user';
import { getCookieValueService, serviceSetCookie } from '@/utils/cookieUtils';
import { getUserFromToken } from '@/utils/jwtUtils';
import dayjs from 'dayjs';
import React, { createContext, useState, ReactNode, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { confirmActionAlert } from '@/utils/alertUtils';
import { signIn } from '@/services/authService';

export interface AuthContextType {
  user: User | null;
  login: (formData: LoginForm) => Promise<void>;
  logout: () => void;
  loading: boolean;
  expireSession: number;
  handleUpdateExpireSession: () => void;
  loadingLogout: boolean;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {

  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [expireSession, setExpireSession] = useState<number>(0);
  const [loadingLogout, setLoadingLogout] = useState<boolean>(false);

  useEffect(() => {
    validateAuth()
  }, [])


  const validateAuth = async () => {

    const token = getCookieValueService(JWT_COOKIE_NAME);

    setLoading(true);

    if (token) {
      const decodedToken = getUserFromToken(token);
      setUser(decodedToken as User);
      const seconds = dayjs.unix(decodedToken.exp as number).diff(dayjs(), 'second');
      serviceSetCookie(JWT_COOKIE_NAME, token, seconds);
      handleUpdateExpireSession();
    } else {
      setUser(null);
    }

    setLoading(false);

  };

  const login = async (formData: LoginForm) => {
    try {
      const { token } = await signIn(formData);
      const decodedToken = getUserFromToken(token);
      setUser(decodedToken as User);
      const seconds = dayjs.unix(decodedToken.exp as number).diff(dayjs(), 'second');
      serviceSetCookie(JWT_COOKIE_NAME, token, seconds);
      handleUpdateExpireSession();
    } catch (error) {
      throw error;
    }
  };

  const logout = async () => {
    const confirm = await confirmActionAlert('Cerrar sesión', '¿Esta seguro de realizar esta acción?', 'question')

    if (confirm) {
      setLoadingLogout(true);
      removeSession();
      await new Promise(resolve => setTimeout(resolve, 2000));
      setLoadingLogout(false);
    }
  };

  const removeSession = () => {
    setUser(null);
    router.push('/login');
    document.cookie = `${JWT_COOKIE_NAME}=; max-age=0`;
  };

  const handleUpdateExpireSession = () => {
    const decodedToken = getUserFromToken(getCookieValueService(JWT_COOKIE_NAME) || "");
    setExpireSession(decodedToken.exp);
  };

  return (
    <AuthContext.Provider value={{
      user,
      login,
      logout,
      loading,
      expireSession,
      handleUpdateExpireSession,
      loadingLogout,
    }}>
      {children}
    </AuthContext.Provider>
  );
};
