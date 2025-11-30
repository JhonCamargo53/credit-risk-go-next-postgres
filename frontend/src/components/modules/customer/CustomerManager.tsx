import GenericTable from '@/components/common/tables/GenericTable';
import { DataColumn } from '@/types/table';
import React, { useEffect, useState } from 'react'
import Button from '@/components/common/buttons/Button';
import { FaPlus } from 'react-icons/fa';
import GenericModal from '@/components/common/modals/GenericModal';
import CustomerRow from './rows/CustomerRow';
import CreateUpdateCustomerForm from './forms/CreateUpdateCustomerForm';
import { useDocumentTypeStore } from '@/store/useDocumentTypeStore';
import { useCustomer } from '@/hooks/useCustomer';
import { generateAxiosErrorToast } from '@/utils/toastUtils';

const userColumns: readonly DataColumn[] = [
    { id: 'index', label: 'N°' },
    { id: 'fullname', label: 'Nombre completo' },
    { id: 'document', label: 'Identificación', minWidth: 150 },
    { id: 'email', label: 'Correo Electrónico', minWidth: 150 },
    { id: 'details', label: 'Detalles', minWidth: 100 },
    { id: 'creditRequestHistory', label: 'Historial de creditos' },
    { id: 'edit', label: 'Editar' },
    { id: 'delte', label: 'Eliminar' }
];

const CustomerManager = () => {

    const { fetchDocumentTypes } = useDocumentTypeStore()
    const { customers, loading, fetchCustomers, handlePagination, pageData } = useCustomer();
    const [createUpdateModal, setCreateUpdateModal] = useState(false);

    useEffect(() => {
        loadCustomers();
    }, []);

    const loadCustomers = async () => {
        try {
            await fetchDocumentTypes()
            await fetchCustomers()
        } catch (error) {
            generateAxiosErrorToast(error, 'Error al cargar los clientes', 'Inténtalo nuevamente');
        }
    }

    return (<div className="space-y-4">

        <div className="grid grid-cols-12 gap-4">
            <div className="col-span-2">
                <Button iconLeft={<FaPlus />} className="font-bold w-full h-full" onClick={() => setCreateUpdateModal(true)}>Agregar</Button>
            </div>
        </div>

        <div>
            <GenericTable
                columns={userColumns}
                data={customers}
                row={<CustomerRow />}
                handlePagination={handlePagination}
                loading={loading.fetching}
                pageData={pageData}
            />
        </div>

        {createUpdateModal && <GenericModal content={<CreateUpdateCustomerForm />} isOpen={createUpdateModal} setOpen={setCreateUpdateModal} size="2xl" />}

    </div>

    )
}

export default CustomerManager