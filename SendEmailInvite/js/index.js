const { Configuration, DefaultApi } = require('userup-sdk')
const nodemailer = require('nodemailer')
require('dotenv').config()

async function main() {
    const smtpPassword = process.env.USERSERVICE_SMTP_PASSWORD
    const smtpUsername = process.env.USERSERVICE_SMTP_SENDEREMAIL
    const email = process.env.USERSERVICE_EMAIL

    const API_BASE_PATH = 'http://localhost:9001/api'

    const config = new Configuration({basePath: API_BASE_PATH})
    const api = new DefaultApi(config)

    try {
        const user = await api.createUser({
            user: {
                username: email,
                attributes: {
                    status: 'pending',
                },
            }
        })

        api.createEvent({
            id: user.iD,
            event: {
                specversion: '1.0',
                type: 'InviteUserEvent',
                time: new Date(),
                datacontenttype: 'application/json',
                subject: user.iD.toString(),
                id: '123',
                data: user
            }
        })

        await sendEmail(email, smtpUsername, smtpPassword, 'Invite User', 'You are invited to userup.io')
    } catch (error) {
        if (error.response) {
            console.error(`Error: ${error.message}`);
            console.error(`Status: ${error.response.status}`);
            console.error(`Headers: ${JSON.stringify(error.response.headers, null, 2)}`);
            console.error(`Body: ${JSON.stringify(error.response.data, null, 2)}`);
        } else {
            console.error(`Error: ${error.message}`);
        }
    }
}

async function sendEmail(toEmail, username, password, subject, body) {
    let transporter = nodemailer.createTransport({
        host: 'smtp.gmail.com',
        port: 587,
        secure: false,
        auth: {
            user: username,
            pass: password,
        },
    })

    let info = await transporter.sendMail({
        from: username,
        to: toEmail,
        subject: subject,
        text: body,
    })

    console.log('Message sent: %s', info.messageId)
}

main().catch(console.error)
