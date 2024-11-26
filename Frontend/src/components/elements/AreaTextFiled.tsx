import { TextField, TextFieldProps } from '@mui/material';

const AreaTextField: React.FC<TextFieldProps> = (props) => {
    return (
        <TextField
            variant="outlined"
            fullWidth
            margin="normal"
            sx={{
                ...props.sx
            }}
            {...props}
        />
    );
};

export { AreaTextField };