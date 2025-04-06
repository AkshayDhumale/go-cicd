# 🔐 Azure Key Vault: Secrets Management Guide

![Azure Key Vault Banner](https://learn.microsoft.com/en-us/azure/key-vault/media/overview-key-vault.png)

This document covers how to create an Azure Key Vault, manage access (users & apps), add secrets, and enforce policies using **Azure Policy**.

## 📘 Table of Contents
* [🔧 Create Azure Key Vault](#create-azure-key-vault)
* [🔑 Grant Access to Users and Applications](#grant-access-to-users-and-applications)
   * [🛡️ Using Access Policies](#using-access-policies)
   * [👥 Using Azure RBAC](#using-azure-rbac)
* [📥 Add Secrets to Key Vault](#add-secrets-to-key-vault)
* [📏 Policy Enforcement with Azure Policy](#policy-enforcement-with-azure-policy)

## 🔧 Create Azure Key Vault

![Create Key Vault](https://learn.microsoft.com/en-us/azure/key-vault/general/media/quick-create-portal/create-a-vault-basics.png)

1. Go to the Azure Portal.
2. Click on ➕ **Create a resource**.
3. Search for **Key Vault** and select it.
4. Fill in the details:
   * 🔤 **Name**: Unique vault name.
   * 🧾 **Subscription**: Select your subscription.
   * 📁 **Resource Group**: Choose or create a resource group.
   * 🌍 **Region**: Select your nearest region.
5. Click **Review + create**, then click **Create**.

![Key Vault Architecture](https://learn.microsoft.com/en-us/azure/key-vault/media/key-vault-ovw-how-to-use.png)

🔗 [Official Docs: Quick Create - Portal](https://learn.microsoft.com/en-us/azure/key-vault/general/quick-create-portal)

## 🔑 Grant Access to Users and Applications

Access can be managed via **Access Policies** or **Azure RBAC**.

### 🛡️ Using Access Policies

![Access Policies](https://learn.microsoft.com/en-us/azure/key-vault/general/media/security-worlds/key-vault-access-policy.png)

1. Go to your Key Vault in the portal.
2. Click **Access Policies** → ➕ **Add Access Policy**.
3. Select:
   * 🔐 **Key/Secret/Certificate Permissions**
4. Click **Select principal** → Choose the user or app.
5. Click **Add**, then **Save**.

🔗 [Docs: Assign Access Policy](https://learn.microsoft.com/en-us/azure/key-vault/general/assign-access-policy-portal)

### 👥 Using Azure RBAC

![RBAC Access](https://learn.microsoft.com/en-us/azure/key-vault/general/media/rbac-permissions-key-vault/key-vault-iam-overview.png)

1. Open your Key Vault → Go to **Access Control (IAM)**.
2. Click ➕ **Add role assignment**.
3. Choose a role like:
   * `Key Vault Reader`
   * `Key Vault Secrets Officer`
4. Assign to **User / Group / Service Principal**.
5. Click **Review + assign**.

🔗 [Docs: RBAC Guide](https://learn.microsoft.com/en-us/azure/key-vault/general/rbac-guide)

## 📥 Add Secrets to Key Vault

![Add Secret](https://learn.microsoft.com/en-us/azure/key-vault/secrets/media/quick-create-portal/add-a-secret.png)

1. Go to your Key Vault → **Secrets** → ➕ **Generate/Import**.
2. Choose **Upload Options**: `Manual`.
3. Enter:
   * 🏷️ **Name**
   * 🔐 **Value**
4. Click **Create**.

![Secret Lifecycle](https://learn.microsoft.com/en-us/azure/key-vault/media/key-vault-ovw-security-worlds.png)

🔗 [Docs: Add Secret - Portal](https://learn.microsoft.com/en-us/azure/key-vault/secrets/quick-create-portal)

## 📏 Policy Enforcement with Azure Policy

![Azure Policy](https://learn.microsoft.com/en-us/azure/governance/policy/media/assign-policy-definition/policy-overview.png)

Use Azure Policy to enforce configurations and standards:

🔍 **Use cases**:
* Audit Key Vault configurations.
* Enforce access/networking/security rules.

🪄 Steps:
1. Go to **Azure Policy** in portal.
2. Browse for built-in Key Vault policy definitions.
3. Click **Assign** → Choose scope (subscription/resource group).

![Policy Enforcement](https://learn.microsoft.com/en-us/azure/governance/policy/media/assign-policy-definition/select-available-definition.png)

🔗 [Docs: Azure Policy + Key Vault](https://learn.microsoft.com/en-us/azure/key-vault/general/azure-policy)
