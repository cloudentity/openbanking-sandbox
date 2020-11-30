import React, {useState} from "react";
import {useHistory, useParams} from 'react-router';
import {Breadcrumbs, Link, Snackbar, SnackbarContent, Typography} from "@material-ui/core";

export default () => {
    const {id} = useParams<{ id: string }>();
    const [snackbar, setSnackbar] = useState<any>(null)
    const history = useHistory();

    const handleBackToAccounts = ()  => {
        history.push('/');
    }

    return (
        <>
            <Breadcrumbs style={{marginBottom: 32}}>
                <Link color="inherit" onClick={handleBackToAccounts}>
                    Accounts
                </Link>
                <Typography color="textPrimary">{id}</Typography>
            </Breadcrumbs>
            <div style={{display: 'flex', alignItems: 'center', justifyContent: 'space-between'}}>
                <Typography variant={'h4'}>Consents list</Typography>
            </div>

            {snackbar && (
                <Snackbar
                    open={true}
                    anchorOrigin={{vertical: 'bottom', horizontal: 'left'}}
                    autoHideDuration={4000}
                    onClose={() => setSnackbar(false)}
                >
                    <SnackbarContent
                        style={{background: snackbar.background}}
                        aria-describedby="client-snackbar"
                        message={<span>{snackbar.message}</span>}
                    />
                </Snackbar>
            )}
        </>
    )
};
