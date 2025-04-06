# 🔐 Azure Key Vault: Secrets Management Guide

![Azure Key Vault Architecture](https://learn.microsoft.com/en-us/azure/key-vault/media/overview/azure-key-vault-overview.png)  
*Azure Key Vault security model and components :cite[4]:cite[6]*

---

## 🔧 Create Azure Key Vault

![Key Vault Creation Portal](https://learn.microsoft.com/en-us/azure/key-vault/media/quick-create-portal/portal-create-resource.png)  
*Step 2-3: Creating a Key Vault in Azure Portal :cite[1]*

**Best Practice:**  
Create separate vaults per application environment (dev/test/prod) :cite[5]. Always enable purge protection during creation.

---

## 🔑 Grant Access to Users and Applications

### 🛡️ Access Policies vs RBAC
![Access Control Comparison](https://learn.microsoft.com/en-us/azure/key-vault/media/rbac-guide/access-control-methods.png)  
*Choosing between legacy access policies and modern RBAC :cite[4]*

**Zero Trust Tip:**  
Implement JIT (Just-In-Time) access through PIM for privileged operations :cite[3]:cite[6].

---

## 📥 Add Secrets to Key Vault

![Secret Creation Interface](https://learn.microsoft.com/en-us/azure/key-vault/media/quick-create-portal/secret-import.png)  
*Secure secret upload with manual entry :cite[1]*

**Security Recommendation:**  
- Rotate secrets every 60-90 days :cite[5]
- Store complex credentials as JSON objects  
`{"username": "admin", "password": "P@ssw0rd!123"}` :cite[8]

---

## 📏 Policy Enforcement with Azure Policy

![Key Vault Policy Dashboard](https://learn.microsoft.com/en-us/azure/governance/policy/media/assign-policy-portal/policy-assignment-details.png)  
*Enforcing encryption standards and network restrictions :cite[4]:cite[5]*

**Essential Policies:**  
1. Audit Vaults without firewall rules  
2. Enforce TLS 1.2+ compliance  
3. Mandate expiration dates for secrets :cite[5]

---

## 🛡️ Advanced Security Setup

```mermaid
graph TD
    A[Application] -->|MSI| B(Key Vault)
    B --> C[HSM-Backed Encryption]
    C --> D[Azure Monitor]
    D --> E[Security Alerts]
