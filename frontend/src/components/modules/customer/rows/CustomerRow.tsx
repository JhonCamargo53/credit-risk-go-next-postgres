import Button from '@/components/common/buttons/Button';
import GenericModal from '@/components/common/modals/GenericModal';
import { GenericTableRow } from '@/types/table'
import React, { useState } from 'react'
import { generateAxiosErrorToast, showSuccessToast } from '@/utils/toastUtils';
import { FaEye, FaPencilAlt } from 'react-icons/fa';
import CreateUpdateUserForm from '../forms/CreateUpdateCustomerForm';
import { Customer } from '@/types/customer';
import CustomerDetailsCard from '../cards/CustomerDetailsCard';
import { GiReceiveMoney } from "react-icons/gi";
import CreditRequestManager from '../../credit-request/CreditRequestManager';
import { RiDeleteBin6Fill } from 'react-icons/ri';
import { confirmActionAlert } from '@/utils/alertUtils';
import { useCustomer } from '@/hooks/useCustomer';
import { useDocumentTypeStore } from '@/store/useDocumentTypeStore';

interface CustomerRow extends GenericTableRow<Customer> {

}

const CustomerRow: React.FC<CustomerRow> = ({ index = 0, data, pageData = { currentPage: 1, limit: 10, lastPage: -1 } }) => {

    const [customer] = useState(data!);

    const [customerDetailsModal, setCustomerDetailsModal] = useState(false);
    const [customerUpdateModal, setCustomerUpdateModal] = useState(false);
    const [creditRequestHistoryModal, setCreditRequestHistoryModal] = useState(false);

    const [deleting, setDeleting] = useState(false);
    const { deleteCustomer } = useCustomer()
    const {documentTypes} = useDocumentTypeStore()
    
    const documentType = documentTypes.find(dt => dt.ID === customer.documentTypeId)?.code || customer.documentTypeId;

    const handleDelete = async () => {
        try {

            const confirm = await confirmActionAlert('Eliminar Cliente', '¿Está seguro de realizar esta acción?', 'question')

            if (confirm) {
                setDeleting(true)
                await deleteCustomer(customer.ID);
                showSuccessToast('Eliminado', 'El  cliente ha sido eliminado correctamente.');
            }

        } catch (error) {
            generateAxiosErrorToast(error, 'Error al eliminar', 'Inténtalo nuevamente');

        } finally {
            setDeleting(false);
        }
    }

    return (
        <tr key={customer.ID} className={`border-b border-gray-300 hover:bg-gray-200 text-center font-medium ${index % 2 !== 0 ? 'bg-primary/8' : ''}`}>
            <td className="px-4 py-2">{index + 1 + (pageData.currentPage - 1) * pageData.limit}</td>
            <td className="px-4 py-2 whitespace-nowrap">{customer.name}</td>
            <td className="px-4 py-2 whitespace-nowrap">{documentType} - {customer.documentNumber}</td>
            <td className="px-4 py-2">{customer.email}</td>
            <td className="px-4 py-2">
                <div className="flex items-center justify-center h-full">
                    <Button title="Detalles Usuario" onClick={() => setCustomerDetailsModal(true)}>
                        <FaEye />
                    </Button>
                </div>
            </td>

            <td className="px-4 py-2 ">
                <div className="flex items-center justify-center h-full">
                    <Button title="Historial de creditos" onClick={() => setCreditRequestHistoryModal(true)}>
                        <GiReceiveMoney />
                    </Button>
                </div>
            </td>
            <td className="px-4 py-2 ">
                <div className="flex items-center justify-center h-full">
                    <Button title="Editar usuario" onClick={() => setCustomerUpdateModal(true)}>
                        <FaPencilAlt />
                    </Button>
                </div>
            </td>
            <td className="px-4 py-2 ">
                <div className="flex items-center justify-center h-full">
                    <Button title="Eliminar cliente" loading={deleting} onClick={() => handleDelete()}>
                        <RiDeleteBin6Fill />
                    </Button>
                </div>
            </td>

            {customerDetailsModal &&
                <GenericModal content={<CustomerDetailsCard customer={customer} />} isOpen={customerDetailsModal}
                    setOpen={setCustomerDetailsModal} size='2xl'
                />}

            {customerUpdateModal &&
                <GenericModal content={<CreateUpdateUserForm customer={customer} />}
                    isOpen={customerUpdateModal} setOpen={setCustomerUpdateModal} size='2xl' />}

            {creditRequestHistoryModal &&
                <GenericModal content={<CreditRequestManager customerId={customer.ID} />}
                    isOpen={creditRequestHistoryModal} setOpen={setCreditRequestHistoryModal} size='5xl' />}

        </tr>
    )
}

export default CustomerRow