'use client';
import React, { createContext, useContext, useState, ReactNode } from 'react';
import { createUser, deleteUser, fetchUserById, fetchUsers, updateUser } from '@/services/userService';
import { User, UserForm } from '@/types/user';
import { PageData } from '@/types/table';

export interface UserContextType {
  users: User[];
  allUsers: User[];
  selectedUser: User | null;

  pageData: PageData;

  loading: {
    fetching: boolean;
    creating: boolean;
    updating: boolean;
    deleting: boolean;
  };

  error: string | null;
  success: string | null;

  resetStatus: () => void;

  handlePagination: (pageData: PageData) => void;

  fetchUsers: () => Promise<void>;
  fetchUserById: (id: number) => Promise<void>;
  createUser: (data: UserForm) => Promise<void>;
  updateUser: (id: number, data: Partial<User>) => Promise<void>;
  deleteUser: (id: number) => Promise<void>;
}

export const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider = ({ children }: { children: ReactNode }) => {

  const [users, setUsers] = useState<User[]>([]);
  const [allUsers, setAllUsers] = useState<User[]>([]);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);

  const [pageData, setPageData] = useState<PageData>({
    currentPage: 1,
    limit: 10,
    total: 0,
    lastPage: 1,
  });

  const [loading, setLoading] = useState({
    fetching: true,
    creating: false,
    updating: false,
    deleting: false,
  });

  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const resetStatus = () => {
    setError(null);
    setSuccess(null);
  };

  const handlePagination = (data: PageData, sourceUsers?: User[]) => {
    const { currentPage, limit } = data;
    const all = sourceUsers || allUsers;

    const start = (currentPage - 1) * limit;
    const end = start + limit;

    const paginated = all.slice(start, end);

    setUsers(paginated);
    setPageData({
      currentPage,
      limit,
      total: all.length,
      lastPage: Math.ceil(all.length / limit),
    });
  };

  const fetchUsersHandler = async () => {
    try {
      setLoading((loading) => ({ ...loading, fetching: true }));
      setError(null);

      const data = await fetchUsers();
      setAllUsers(data);
      handlePagination({ ...pageData, total: data.length }, data);

    } catch (err: any) {
      setError(err.message || 'Error al obtener usuarios');
      throw err;
    } finally {
      setLoading((loading) => ({ ...loading, fetching: false }));
    }
  };

  const fetchUserByIdHandler = async (id: number) => {
    try {
      setLoading((loading) => ({ ...loading, fetching: true }));
      setError(null);

      const data = await fetchUserById(id);
      setSelectedUser(data);

    } catch (err: any) {
      setError(err.message || 'Error al obtener usuario');
      throw err;
    } finally {
      setLoading((loading) => ({ ...loading, fetching: false }));
    }
  };

  const createUserHandler = async (data:UserForm) => {
    try {
      setLoading((loading) => ({ ...loading, creating: true }));
      setError(null);
      setSuccess(null);

      const newUser = await createUser(data);

      const updated = [...allUsers, newUser];
      setAllUsers(updated);
      handlePagination(pageData);

      setSuccess('Usuario creado correctamente');

    } catch (err: any) {
      setError(err.message || 'Error al crear usuario');
      throw err;
    } finally {
      setLoading((loading) => ({ ...loading, creating: false }));
    }
  };

  const updateUserHandler = async (id: number, data: Partial<User>) => {
    try {
      setLoading((loading) => ({ ...loading, updating: true }));
      setError(null);
      setSuccess(null);

      const updatedUser = await updateUser(id, data);

      setAllUsers(prev => {
        const newData = prev.some(c => c.ID === id)
          ? prev.map(user => (user.ID === id ? updatedUser : user))
          : [updatedUser, ...prev];

        handlePagination(pageData, newData);
        return newData;
      });

      setSelectedUser(updatedUser);

      setSuccess('Usuario actualizado correctamente');

    } catch (err: any) {
      setError(err.message || 'Error al actualizar usuario');
      throw err;
    } finally {
      setLoading((loading) => ({ ...loading, updating: false }));
    }
  };

  const deleteUserHandler = async (id: number) => {
    try {
      setLoading((loading) => ({ ...loading, deleting: true }));
      setError(null);
      setSuccess(null);

      await deleteUser(id);

      const updated = allUsers.filter((u) => u.ID !== id);
      setAllUsers(updated);
      1
      handlePagination(pageData);
      setSuccess('Usuario eliminado correctamente');

    } catch (err: any) {
      setError(err.message || 'Error al eliminar usuario');
      throw err;
    } finally {
      setLoading((loading) => ({ ...loading, deleting: false }));
    }
  };

  return (
    <UserContext.Provider
      value={{
        users,
        allUsers,
        selectedUser,
        pageData,
        loading,
        error,
        success,
        resetStatus,
        handlePagination,
        fetchUsers: fetchUsersHandler,
        fetchUserById: fetchUserByIdHandler,
        createUser: createUserHandler,
        updateUser: updateUserHandler,
        deleteUser: deleteUserHandler,
      }}
    >
      {children}
    </UserContext.Provider>
  );
};