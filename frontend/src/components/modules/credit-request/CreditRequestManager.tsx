import Button from '@/components/common/buttons/Button';
import GenericModal from '@/components/common/modals/GenericModal';
import GenericTable from '@/components/common/tables/GenericTable'
import { useCreditRequest } from '@/hooks/useCreditRequest';
import { DataColumn } from '@/types/table';
import React, { useEffect, useState } from 'react'
import { FaPlus } from 'react-icons/fa';
import CreditRequestRow from './rows/CreditRequestRow';
import { useCreditStatusStore } from '@/store/useCreditStatusStore';
import CreateUpdateCreditRequestForm from './form/CreateUpdateCreditRequestForm';
import { generateAxiosErrorToast } from '@/utils/toastUtils';

interface CreditRequestManagerProps {
    customerId?: number
}

const userColumns: readonly DataColumn[] = [
    { id: 'index', label: 'N°' },
    { id: 'amount', label: 'Cantidad' },
    { id: 'termMonths', label: 'Número de meses', minWidth: 150 },
    { id: 'productType', label: 'Proposito', minWidth: 150 },
    { id: 'creditStatus', label: 'Estado', minWidth: 100 },
    { id: 'riskScore', label: 'Nivel de Riesgo' },
    { id: 'details', label: 'Detalles' },
    { id: 'customerAssets', label: 'Ver activos' },
    { id: 'edit', label: 'Editar' },
    { id: 'delete', label: 'Eliminar' }
];


const CreditRequestManager: React.FC<CreditRequestManagerProps> = ({ customerId }) => {

    const [createUpdateModal, setCreateUpdateModal] = useState(false);
    const { fetchCreditStatuses } = useCreditStatusStore()
    const { creditRequests, loading, pageData, fetchCreditRequestsByCustomerId, handlePagination } = useCreditRequest();

    useEffect(() => {
        loadCreditRequests()
    }, [])

    const loadCreditRequests = async () => {
        try {
            await fetchCreditStatuses()
            if (customerId) {
                await fetchCreditRequestsByCustomerId(customerId)
            }
        } catch (error) {
            generateAxiosErrorToast(error, 'Error al cargar las solicitudes de creditos', 'Inténtalo nuevamente');
        }
    }

    return (
        <div className="space-y-4">
            <div className="grid grid-cols-12 gap-4">
                <div className="col-span-5 md:col-span-3">
                    <Button iconLeft={<FaPlus />} className="font-bold w-full h-full" onClick={() => setCreateUpdateModal(true)}>Agregar</Button>
                </div>
            </div>
            <div>
                <GenericTable
                    columns={userColumns}
                    data={creditRequests}
                    row={<CreditRequestRow />}
                    handlePagination={handlePagination}
                    loading={loading.fetching}
                    pageData={pageData}
                />
            </div>
            {createUpdateModal && <GenericModal content={<CreateUpdateCreditRequestForm customerId={customerId} />} isOpen={createUpdateModal} setOpen={setCreateUpdateModal} size="2xl" />}
        </div>
    )
}

export default CreditRequestManager