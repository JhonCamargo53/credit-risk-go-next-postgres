import Button from '@/components/common/buttons/Button';
import GenericModal from '@/components/common/modals/GenericModal';
import { GenericTableRow } from '@/types/table'
import React, { useState } from 'react'
import { FaEye, FaPencilAlt } from 'react-icons/fa';
import CreditRequestManager from '../CreditRequestManager';
import CreditRequestCard from '../cards/CreditRequestCard';
import { useCreditStatusStore } from '@/store/useCreditStatusStore';
import CreateUpdateCreditRequestForm from '../form/CreateUpdateCreditRequestForm';
import { PiBankFill } from "react-icons/pi";
import CustomerAssetManager from '../../customer-asset/CustomerAssetManager';
import { RiDeleteBin6Fill } from 'react-icons/ri';
import { useCreditRequest } from '@/hooks/useCreditRequest';
import { confirmActionAlert } from '@/utils/alertUtils';
import { generateAxiosErrorToast, showSuccessToast } from '@/utils/toastUtils';

interface CreditRequestRow extends GenericTableRow<CreditRequest> {

}

const CreditRequestRow: React.FC<CreditRequestRow> = ({ index = 0, data, pageData = { currentPage: 1, limit: 10, lastPage: -1 } }) => {

    const { creditStatuses } = useCreditStatusStore();

    const [creditRequest] = useState(data!);

    const [customerDetailsModal, setCreditRequestDetailsModal] = useState(false);
    const [customerUpdateModal, setCreditRequestUpdateModal] = useState(false);
    const [creditRequestHistoryModal, setCreditRequestHistoryModal] = useState(false);
    const [customerAssetsModal, setCustomerAssetsModalModal] = useState(false);

    const creditStatus = creditStatuses.find(status => status.ID === creditRequest.creditStatusId);

    const [deleting, setDeleting] = useState(false);
    const { deleteCreditRequest } = useCreditRequest()

    const handleDelete = async () => {
        try {
            const confirm = await confirmActionAlert('Eliminar Solicitud de credito', '¿Está seguro de realizar esta acción?', 'question')
            if (confirm) {
                setDeleting(true)
                await deleteCreditRequest(creditRequest.ID);
                showSuccessToast('Eliminado', 'La solicitud de credito ha sido eliminada correctamente.');
            }
        } catch (error) {
            generateAxiosErrorToast(error, 'Error al eliminar', 'Inténtalo nuevamente');
        } finally {
            setDeleting(false);
        }
    }


    return (
        <tr key={creditRequest.ID} className={`border-b border-gray-300 hover:bg-gray-200 text-center font-medium ${index % 2 !== 0 ? 'bg-primary/8' : ''}`}>
            <td className="px-2 py-2">{index + 1 + (pageData.currentPage - 1) * pageData.limit}</td>
            <td className="px-2 py-2 whitespace-nowrap">{creditRequest.amount.toLocaleString()}</td>
            <td className="px-2 py-2 whitespace-nowrap">{creditRequest.termMonths}</td>
            <td className="px-2 py-2">{creditRequest.productType}</td>
            <td className="px-2 py-2">{creditStatus?.name}</td>
            <td className="px-2 py-2">{creditRequest.riskScore}</td>

            <td className="px-2 py-2">
                <div className="flex items-center justify-center h-full">
                    <Button title="Detalles cliente" onClick={() => setCreditRequestDetailsModal(true)}>
                        <FaEye />
                    </Button>
                </div>
            </td>
            <td className="px-2 py-2">
                <div className="flex items-center justify-center h-full">
                    <Button title="Ver activos del cliente" onClick={() => setCustomerAssetsModalModal(true)}>
                        <PiBankFill />
                    </Button>
                </div>
            </td>
            <td className="px-2 py-2 ">
                <div className="flex items-center justify-center h-full">
                    <Button title="Editar cliente" onClick={() => setCreditRequestUpdateModal(true)}>
                        <FaPencilAlt />
                    </Button>
                </div>
            </td>
            <td className="px-2 py-2 ">
                <div className="flex items-center justify-center h-full">
                    <Button title="Eliminar cliente" loading={deleting} onClick={() => handleDelete()}>
                        <RiDeleteBin6Fill />
                    </Button>
                </div>
            </td>

            {customerDetailsModal &&
                <GenericModal content={<CreditRequestCard creditRequest={creditRequest} />} isOpen={customerDetailsModal}
                    setOpen={setCreditRequestDetailsModal} size='2xl'
                />}

            {customerAssetsModal &&
                <GenericModal content={<CustomerAssetManager creditRequestId={creditRequest.ID} />} isOpen={customerAssetsModal}
                    setOpen={setCustomerAssetsModalModal} size='5xl'
                />}

            {customerUpdateModal &&
                <GenericModal content={<CreateUpdateCreditRequestForm creditRequest={creditRequest} customerId={creditRequest.customerId} />}
                    isOpen={customerUpdateModal} setOpen={setCreditRequestUpdateModal} size='2xl' />}


            {creditRequestHistoryModal &&
                <GenericModal content={<CreditRequestManager customerId={creditRequest.ID} />}
                    isOpen={creditRequestHistoryModal} setOpen={setCreditRequestHistoryModal} size='7xl' />}

        </tr>
    )
}

export default CreditRequestRow