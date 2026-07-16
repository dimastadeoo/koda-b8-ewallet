# ERD Aplikasi Ewallet

## Berikut merupakan ERD dan cara run aplikasi Ewallet

### Akses Aplikasi
Untuk menjalankan aplikasi bisa pull dari package dan coba jalankan dengan -it supaya aplikasi berjalan di terminal

```sh
docker pull ghcr.io/dimastadeoo/koda-b8-ewallet:latest
docker run -it ghcr.io/dimastadeoo/koda-b8-ewallet:latest
```


```mermaid
erDiagram
    users {
        BIGINT id PK
        VARCHAR pin
        VARCHAR email UK
        VARCHAR hp_number UK
        VARCHAR status_account
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    profiles {
        BIGINT id PK
        BIGINT id_user FK
        VARCHAR nik UK
        VARCHAR name
        TEXT address
        VARCHAR gender
        VARCHAR place_birth
        DATE date_birth
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    sessions {
        BIGINT id PK
        VARCHAR token UK
        VARCHAR status
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    users_logs {
        BIGINT id_user FK
        BIGINT id_session FK
        TEXT activity_detail
        VARCHAR ip_address
        TIMESTAMP created_at
    }

    forgot_pin {
        BIGINT id_sessions FK
        BIGINT id_user FK
        VARCHAR previous_pin
        VARCHAR update_pin
        VARCHAR ip_address
        TIMESTAMP created_at
    }

    wallets {
        BIGINT id PK
        BIGINT id_user FK
        NUMERIC balance
        VARCHAR currency
        VARCHAR status_account
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    transactions {
        BIGINT id PK
        BIGINT id_wallet FK
        VARCHAR transaction_type
        NUMERIC amount
        VARCHAR status
        TEXT description
        TIMESTAMP created_at
        TIMESTAMP completed_at
    }

    transfers {
        BIGINT id PK
        BIGINT trx_id FK
        BIGINT sender_wallet_id FK
        BIGINT receiver_wallet_id FK
        TEXT notes
        TIMESTAMP created_at
    }

    wallet_logs {
        BIGINT id PK
        BIGINT wallet_id FK
        VARCHAR action
        TEXT description
        TIMESTAMP created_at
    }

    users ||--|| profiles : has
    users ||--o{ wallets : owns
    users ||--o{ users_logs : activity
    users ||--o{ forgot_pin : reset_pin

    sessions ||--o{ users_logs : records
    sessions ||--o{ forgot_pin : used_for

    wallets ||--o{ transactions : performs
    wallets ||--o{ wallet_logs : logs

    transactions ||--|| transfers : transfer_detail

    wallets ||--o{ transfers : sender
    wallets ||--o{ transfers : receiver
```

### Screenshot

<table>
    <tr>
        <td>Tampilan Awal Login / Register</td>
        <td>Register Users</td>
        <td>Dashboard Aplikasi Setelah login</td>
    </tr>
    <tr>
        <td><img src="img/Screenshot_2026-07-16_20-19-35.png" alt="Tampilan Awal Ewallet"></td>
        <td><img src="img/Screenshot_2026-07-16_20-21-21.png" alt="Register"></td>
        <td><img src="img/Screenshot_2026-07-16_20-21-37.png" alt="Dashboard"></td>
    </tr>
</table>

<table>
    <tr>
        <td>List User</td>
        <td>Topup Saldo</td>
        <td>Tampilkan Saldo</td>
    </tr>
    <tr>
        <td><img src="img/Screenshot_2026-07-16_20-22-27.png" alt="List Users"></td>
        <td><img src="img/Screenshot_2026-07-16_20-23-05.png" alt="Topup saldo"></td>
        <td><img src="img/Screenshot_2026-07-16_20-25-06.png" alt="Tampilan Saldo"/></td>
    </tr>
</table>

<table>
    <tr>
        <td>Process Transfer</td>
        <td>Topup Saldo</td>
    </tr>
    <tr>
        <td><img src="img/Screenshot_2026-07-16_20-25-25.png" alt="Transfer"></td>
        <td><img src="img/Screenshot_2026-07-16_20-25-39.png" alt="History"></td>
    </tr>
</table>