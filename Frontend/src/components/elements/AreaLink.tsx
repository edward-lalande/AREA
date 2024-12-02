import { Link, LinkProps } from '@mui/material';

interface AreaLinkProps extends LinkProps {
    text: string;
}

const AreaLink: React.FC<AreaLinkProps> = ({ text, ...props }) => {
    return (
        <Link
            {...props}
            sx={{
                color: "black",
                textDecoration: "underline #000000",
                ...props.sx
            }}
		>
            {text}
        </Link>
    );
};
  
export { AreaLink };
