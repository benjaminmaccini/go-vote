# TODO List

### 08-21-2024

- [ ] Reassess design
- [x] Move to sourcehut (with mirror for github)
- [ ] Cleanup tasks after revisiting
  - [ ] Drop dependencies (where it makes sense)
  - [ ] Change license to AGPL or similar
  - [ ] Update logging to be complete
  - [ ] Break out protocol into different files
  - [ ] Add test infrastructure
  - [ ] Remove in-memory global election variable and persist to SQLite

### 08-22-2024

- [ ] Ranked choice
- [ ] Multi member constituency methods
- [ ] Concurrent elections

### Later

- [ ] Add [cryptographic tallying](http://security.hsr.ch/msevote/seminar-papers/HS09_Homomorphic_Tallying_with_Paillier.pdf)
- [ ] Digital Signatures (multisignature?)
- [ ] Verifiable server state (version the same between server-client? Non-malicious server accepting votes?)
