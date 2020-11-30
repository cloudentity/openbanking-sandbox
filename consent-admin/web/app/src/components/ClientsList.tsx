import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import {Button, Typography} from "@material-ui/core";

const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
    withHover: {
        '&:hover': {
            cursor: 'pointer'
        }
    },
    consent: {
        "&:nth-child(odd)": {
            background: '#F4F7F8'
        },
        "&:nth-child(even)": {
            background: '#feffff'
        }
    }
});

export default ({clients, onRevokeClient, onRevokeConsent}) => {
    const classes = useStyles();

    return (
        <>
            <Typography variant={'h4'}>List of TTPs with granted account access consents</Typography>
            <div style={{marginTop: 32}}>
                {clients.map(client => (
                    <Paper style={{marginBottom: 32}}>
                        <div style={{
                            padding: 16,
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'space-between',
                            borderLeft: "6px solid #006580"
                        }}>
                            <div>
                                <Typography>Client Name: <strong>{client.client_name}</strong></Typography>
                                <Typography variant={'caption'} style={{marginTop: 8}}
                                            display={'block'}>Client
                                    ID: <strong>{client.client_id}</strong></Typography>
                            </div>
                            <Button onClick={onRevokeClient(client.client_id)} variant={'outlined'}>Revoke all client consents</Button>
                        </div>
                        <div style={{}}>
                            {client.consents.map(consent => (
                                <div style={{padding: 16}} className={classes.consent}>
                                    <div style={{display: 'flex', alignItems: 'center'}}>
                                        <Typography>Consent ID: </Typography>
                                        <Typography><strong>{consent.consent_id}</strong></Typography>
                                        <div style={{flex: 1}}/>
                                        <Button onClick={onRevokeConsent(consent.consent_id)}>Revoke consent</Button>
                                    </div>
                                    <div style={{display: 'flex', marginTop: 16}}>
                                        <Typography>Permissions: </Typography>
                                        <Typography> <strong>{consent.permissions?.join(', ')}</strong></Typography>
                                    </div>
                                    <div style={{display: 'flex', marginTop: 8}}>
                                        <Typography>Accounts ID: </Typography>
                                        <Typography> <strong>{consent.account_ids?.join(', ')}</strong></Typography>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </Paper>
                ))}
            </div>

            {/*<TableContainer component={Paper} style={{marginTop: 32}}>*/}
            {/*  <Table className={classes.table} aria-label="Accounts table">*/}
            {/*    <TableHead>*/}
            {/*      <TableRow>*/}
            {/*        <TableCell>Name</TableCell>*/}
            {/*        <TableCell align="right">ID</TableCell>*/}
            {/*        <TableCell align="right"></TableCell>*/}
            {/*      </TableRow>*/}
            {/*    </TableHead>*/}
            {/*    <TableBody>*/}
            {/*      {clients.map((row) => (*/}
            {/*        <TableRow key={row.client_id}>*/}
            {/*          <TableCell component="th" scope="row">*/}
            {/*            {row.client_name}*/}
            {/*          </TableCell>*/}
            {/*          <TableCell align="right">{row.client_id}</TableCell>*/}
            {/*          <TableCell align="right"><Button onClick={onRevokeClient(row)}>Revoke</Button></TableCell>*/}
            {/*        </TableRow>*/}
            {/*      ))}*/}
            {/*    </TableBody>*/}
            {/*  </Table>*/}
            {/*</TableContainer>*/}
        </>
    );
}
