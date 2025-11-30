'use client';

import React, { createContext, useState, ReactNode, useEffect } from 'react';
import { PageData } from '@/types/table';
import { Customer, CustomerForm } from '@/types/customer';
import { createCustomer, deleteCustomer, fetchCustomerById, fetchCustomers, updateCustomer } from '@/services/customerService';

export interface CustomerContextType {
    customers: Customer[];
    allCustomers: Customer[];
    selectedCustomer: Customer | null;

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

    fetchCustomers: () => Promise<void>;
    fetchCustomerById: (id: number) => Promise<void>;
    createCustomer: (data: CustomerForm) => Promise<void>;
    updateCustomer: (id: number, data: Partial<Customer>) => Promise<void>;
    deleteCustomer: (id: number) => Promise<void>;
}

export const CustomerContext = createContext<CustomerContextType | undefined>(undefined);

export const CustomerProvider = ({ children }: { children: ReactNode }) => {
    const [customers, setCustomers] = useState<Customer[]>([]);
    const [allCustomers, setAllCustomers] = useState<Customer[]>([]);
    const [selectedCustomer, setSelectedCustomer] = useState<Customer | null>(null);

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

    const handlePagination = (data: PageData, sourceCustomers?: Customer[]) => {
        const { currentPage, limit } = data;
        const all = sourceCustomers || allCustomers;

        const start = (currentPage - 1) * limit;
        const end = start + limit;

        const paginated = all.slice(start, end);

        setCustomers(paginated);
        setPageData({
            currentPage,
            limit,
            total: all.length,
            lastPage: Math.ceil(all.length / limit),
        });
    };

    const fetchCustomersAction = async () => {
        try {
            setLoading(prev => ({ ...prev, fetching: true }));
            setError(null);

            const data = await fetchCustomers();
            setAllCustomers(data);

            handlePagination({ ...pageData, total: data.length }, data);
        } catch (err: any) {
            setError(err.message || 'Error al obtener clientes');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, fetching: false }));
        }
    };

    const fetchCustomerByIdAction = async (id: number) => {
        try {
            setLoading(prev => ({ ...prev, fetching: true }));
            setError(null);

            const data = await fetchCustomerById(id);
            setSelectedCustomer(data);
        } catch (err: any) {
            setError(err.message || 'Error al obtener cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, fetching: false }));
        }
    };

    const createCustomerAction = async (data: CustomerForm) => {
        try {
            setLoading(prev => ({ ...prev, creating: true }));
            setError(null);
            setSuccess(null);

            const newCustomer = await createCustomer(data);
            setAllCustomers(prev => {
                const newData = [newCustomer, ...prev];
                handlePagination(pageData, newData);
                return newData;
            });

            handlePagination(pageData);
            setSuccess('Cliente creado correctamente');
        } catch (err: any) {
            setError(err.message || 'Error al crear cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, creating: false }));
        }
    };

    const updateCustomerAction = async (id: number, data: Partial<Customer>) => {
        try {
            setLoading(prev => ({ ...prev, updating: true }));
            setError(null);
            setSuccess(null);

            const updatedCustomer = await updateCustomer(id, data);

            setAllCustomers(prev => {
                const newData = prev.some(client => client.ID === id)
                    ? prev.map(client => (client.ID === id ? updatedCustomer : client))
                    : [updatedCustomer, ...prev];
                handlePagination(pageData, newData);
                return newData;
            });

            setSelectedCustomer(updatedCustomer);

            setSuccess('Cliente actualizado correctamente');

        } catch (err: any) {
            setError(err.message || 'Error al actualizar cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, updating: false }));
        }
    };

    const deleteCustomerAction = async (id: number) => {
        try {
            setLoading(prev => ({ ...prev, deleting: true }));
            setError(null);
            setSuccess(null);

            await deleteCustomer(id);
            const updatedAll = allCustomers.filter(c => c.ID !== id);
            setAllCustomers(updatedAll);

            handlePagination(pageData,updatedAll);
            setSuccess('Cliente eliminado correctamente');
        } catch (err: any) {
            setError(err.message || 'Error al eliminar cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, deleting: false }));
        }
    };

    return (
        <CustomerContext.Provider value={{
            customers,
            allCustomers,
            selectedCustomer,
            pageData,
            loading,
            error,
            success,
            resetStatus,
            handlePagination,
            fetchCustomers: fetchCustomersAction,
            fetchCustomerById: fetchCustomerByIdAction,
            createCustomer: createCustomerAction,
            updateCustomer: updateCustomerAction,
            deleteCustomer: deleteCustomerAction,
        }}>
            {children}
        </CustomerContext.Provider>
    );
};

