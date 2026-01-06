import axios from 'axios';

const api = axios.create({
    baseURL: '/api', // Proxied by Vite to localhost:8080
    headers: {
        'Content-Type': 'application/json',
    },
});

export const authService = {
    login: (data) => api.post('/users/_Login', data),
    logout: () => api.delete('/users'),
    getCurrentUser: () => api.get('/users/_current'),
    updateProfile: (data) => api.patch('/users/_current', data),
};

export const adminService = {
    checkAccess: () => api.get('/admin/'),
    // Contacts
    getContacts: () => api.get('/admin/contacts'),
    createContact: (data) => api.post('/admin/contacts', data),
    getContact: (id) => api.get(`/admin/contacts/${id}`),
    updateContact: (id, data) => api.put(`/admin/contacts/${id}`, data),
    deleteContact: (id) => api.delete(`/admin/contacts/${id}`),

    // Addresses
    getAddresses: (contactId) => api.get(`/admin/contacts/${contactId}/addresses`),
    createAddress: (contactId, data) => api.post(`/admin/contacts/${contactId}/addresses`, data),
    getAddress: (contactId, addressId) => api.get(`/admin/contacts/${contactId}/addresses/${addressId}`),
    updateAddress: (contactId, addressId, data) => api.put(`/admin/contacts/${contactId}/addresses/${addressId}`, data),
    deleteAddress: (contactId, addressId) => api.delete(`/admin/contacts/${contactId}/addresses/${addressId}`),
};

export const driverService = {
    checkAccess: () => api.get('/driver/'),
};

export const superAdminService = {
    checkAccess: () => api.get('/superadmin/'),
};

export default api;
