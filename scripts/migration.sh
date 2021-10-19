#!/usr/bin/env bash

function postservice_cmd() {
    local service_migration_dir="internal/apps/postservice/db/migrations"
    if [ "$1" == "create" ]; then
        if [ -z "$2" ]; then
            echo "Usage: $0 postservice create <migration name>"
            exit 1
        fi

        echo "Creating migration..."
        migrate create -ext sql -dir "$service_migration_dir" -seq $2
    fi

    if [ "$1" == "up" ]; then
        echo "Migrating up..."
        migrate up -database "$DATABASE_URL" -path "$service_migration_dir" up
    fi

    if [ "$1" == "down" ]; then
        echo "Migrating down..."
        migrate down -database "$DATABASE_URL" -path "$service_migration_dir" down 1
    fi
}

function authservice_cmd() {
    local service_migration_dir="internal/apps/authservice/db/migrations"
    if [ "$1" == "create" ]; then
        if [ -z "$2" ]; then
            echo "Usage: $0 authservice create <migration name>"
            exit 1
        fi

        echo "Creating migration..."
        migrate create -ext sql -dir "$service_migration_dir" -seq $2
    fi

    if [ "$1" == "up" ]; then
        echo "Migrating up..."
        migrate up -database "$DATABASE_URL" -path "$service_migration_dir" up
    fi

    if [ "$1" == "down" ]; then
        echo "Migrating down..."
        migrate down -database "$DATABASE_URL" -path "$service_migration_dir" down 1
    fi
}

function userservice_cmd() {
    local service_migration_dir="internal/apps/userservice/db/migrations"
    if [ "$1" == "create" ]; then
        if [ -z "$2" ]; then
            echo "Usage: $0 userservice create <migration name>"
            exit 1
        fi

        echo "Creating migration..."
        migrate create -ext sql -dir "$service_migration_dir" -seq $2
    fi

    if [ "$1" == "up" ]; then
        echo "Migrating up..."
        migrate up -database "$DATABASE_URL" -path "$service_migration_dir" up
    fi

    if [ "$1" == "down" ]; then
        echo "Migrating down..."
        migrate down -database "$DATABASE_URL" -path "$service_migration_dir" down 1
    fi
}


function root_cmd() {
    # $1: service name [postservice, authservice, userservice]
    # $2: migration command [create, up, down]
    # $3: migration name (only for create)
    if [ -z "$1" ] || [ -z "$2" ]; then
        echo "Usage: $0 <service name> <create|up|down> <migration name>"
        exit 1
    fi

    # postservice
    if [ "$1" == "postservice" ]; then
        postservice_cmd $2 $3
    fi

    # authservice
    if [ "$1" == "authservice" ]; then
        authservice_cmd $2 $3
    fi

    # userservice
    if [ "$1" == "userservice" ]; then
        userservice_cmd $2 $3
    fi
}
root_cmd $1 $2 $3
