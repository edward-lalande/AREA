import React, { useEffect } from 'react';

const Download: React.FC = () => {
    useEffect(() => {
        const link = document.createElement('a');
        link.href = 'README.md'; // Changed by the shared volumes
        link.download = 'area.apk';
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }, []);

    return null;
};

export default Download;
