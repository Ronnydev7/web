// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/emailcredential"
	"api/ent/entgenerated/loginsession"
	"api/ent/entgenerated/predicate"
	"api/ent/entgenerated/user"
	"api/ent/entgenerated/userpublicprofile"
	"errors"
	"fmt"
	"time"
)

// EmailCredentialWhereInput represents a where input for filtering EmailCredential queries.
type EmailCredentialWhereInput struct {
	Predicates []predicate.EmailCredential  `json:"-"`
	Not        *EmailCredentialWhereInput   `json:"not,omitempty"`
	Or         []*EmailCredentialWhereInput `json:"or,omitempty"`
	And        []*EmailCredentialWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "email" field predicates.
	Email             *string  `json:"email,omitempty"`
	EmailNEQ          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGT           *string  `json:"emailGT,omitempty"`
	EmailGTE          *string  `json:"emailGTE,omitempty"`
	EmailLT           *string  `json:"emailLT,omitempty"`
	EmailLTE          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *EmailCredentialWhereInput) AddPredicates(predicates ...predicate.EmailCredential) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the EmailCredentialWhereInput filter on the EmailCredentialQuery builder.
func (i *EmailCredentialWhereInput) Filter(q *EmailCredentialQuery) (*EmailCredentialQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyEmailCredentialWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyEmailCredentialWhereInput is returned in case the EmailCredentialWhereInput is empty.
var ErrEmptyEmailCredentialWhereInput = errors.New("entgenerated: empty predicate EmailCredentialWhereInput")

// P returns a predicate for filtering emailcredentials.
// An error is returned if the input is empty or invalid.
func (i *EmailCredentialWhereInput) P() (predicate.EmailCredential, error) {
	var predicates []predicate.EmailCredential
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, emailcredential.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.EmailCredential, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, emailcredential.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.EmailCredential, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, emailcredential.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, emailcredential.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, emailcredential.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, emailcredential.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, emailcredential.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, emailcredential.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, emailcredential.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, emailcredential.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, emailcredential.IDLTE(*i.IDLTE))
	}
	if i.Email != nil {
		predicates = append(predicates, emailcredential.EmailEQ(*i.Email))
	}
	if i.EmailNEQ != nil {
		predicates = append(predicates, emailcredential.EmailNEQ(*i.EmailNEQ))
	}
	if len(i.EmailIn) > 0 {
		predicates = append(predicates, emailcredential.EmailIn(i.EmailIn...))
	}
	if len(i.EmailNotIn) > 0 {
		predicates = append(predicates, emailcredential.EmailNotIn(i.EmailNotIn...))
	}
	if i.EmailGT != nil {
		predicates = append(predicates, emailcredential.EmailGT(*i.EmailGT))
	}
	if i.EmailGTE != nil {
		predicates = append(predicates, emailcredential.EmailGTE(*i.EmailGTE))
	}
	if i.EmailLT != nil {
		predicates = append(predicates, emailcredential.EmailLT(*i.EmailLT))
	}
	if i.EmailLTE != nil {
		predicates = append(predicates, emailcredential.EmailLTE(*i.EmailLTE))
	}
	if i.EmailContains != nil {
		predicates = append(predicates, emailcredential.EmailContains(*i.EmailContains))
	}
	if i.EmailHasPrefix != nil {
		predicates = append(predicates, emailcredential.EmailHasPrefix(*i.EmailHasPrefix))
	}
	if i.EmailHasSuffix != nil {
		predicates = append(predicates, emailcredential.EmailHasSuffix(*i.EmailHasSuffix))
	}
	if i.EmailEqualFold != nil {
		predicates = append(predicates, emailcredential.EmailEqualFold(*i.EmailEqualFold))
	}
	if i.EmailContainsFold != nil {
		predicates = append(predicates, emailcredential.EmailContainsFold(*i.EmailContainsFold))
	}

	if i.HasOwner != nil {
		p := emailcredential.HasOwner()
		if !*i.HasOwner {
			p = emailcredential.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOwnerWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, emailcredential.HasOwnerWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyEmailCredentialWhereInput
	case 1:
		return predicates[0], nil
	default:
		return emailcredential.And(predicates...), nil
	}
}

// LoginSessionWhereInput represents a where input for filtering LoginSession queries.
type LoginSessionWhereInput struct {
	Predicates []predicate.LoginSession  `json:"-"`
	Not        *LoginSessionWhereInput   `json:"not,omitempty"`
	Or         []*LoginSessionWhereInput `json:"or,omitempty"`
	And        []*LoginSessionWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "last_login_time" field predicates.
	LastLoginTime      *time.Time  `json:"lastLoginTime,omitempty"`
	LastLoginTimeNEQ   *time.Time  `json:"lastLoginTimeNEQ,omitempty"`
	LastLoginTimeIn    []time.Time `json:"lastLoginTimeIn,omitempty"`
	LastLoginTimeNotIn []time.Time `json:"lastLoginTimeNotIn,omitempty"`
	LastLoginTimeGT    *time.Time  `json:"lastLoginTimeGT,omitempty"`
	LastLoginTimeGTE   *time.Time  `json:"lastLoginTimeGTE,omitempty"`
	LastLoginTimeLT    *time.Time  `json:"lastLoginTimeLT,omitempty"`
	LastLoginTimeLTE   *time.Time  `json:"lastLoginTimeLTE,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *LoginSessionWhereInput) AddPredicates(predicates ...predicate.LoginSession) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the LoginSessionWhereInput filter on the LoginSessionQuery builder.
func (i *LoginSessionWhereInput) Filter(q *LoginSessionQuery) (*LoginSessionQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyLoginSessionWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyLoginSessionWhereInput is returned in case the LoginSessionWhereInput is empty.
var ErrEmptyLoginSessionWhereInput = errors.New("entgenerated: empty predicate LoginSessionWhereInput")

// P returns a predicate for filtering loginsessions.
// An error is returned if the input is empty or invalid.
func (i *LoginSessionWhereInput) P() (predicate.LoginSession, error) {
	var predicates []predicate.LoginSession
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, loginsession.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.LoginSession, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, loginsession.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.LoginSession, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, loginsession.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, loginsession.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, loginsession.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, loginsession.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, loginsession.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, loginsession.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, loginsession.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, loginsession.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, loginsession.IDLTE(*i.IDLTE))
	}
	if i.LastLoginTime != nil {
		predicates = append(predicates, loginsession.LastLoginTimeEQ(*i.LastLoginTime))
	}
	if i.LastLoginTimeNEQ != nil {
		predicates = append(predicates, loginsession.LastLoginTimeNEQ(*i.LastLoginTimeNEQ))
	}
	if len(i.LastLoginTimeIn) > 0 {
		predicates = append(predicates, loginsession.LastLoginTimeIn(i.LastLoginTimeIn...))
	}
	if len(i.LastLoginTimeNotIn) > 0 {
		predicates = append(predicates, loginsession.LastLoginTimeNotIn(i.LastLoginTimeNotIn...))
	}
	if i.LastLoginTimeGT != nil {
		predicates = append(predicates, loginsession.LastLoginTimeGT(*i.LastLoginTimeGT))
	}
	if i.LastLoginTimeGTE != nil {
		predicates = append(predicates, loginsession.LastLoginTimeGTE(*i.LastLoginTimeGTE))
	}
	if i.LastLoginTimeLT != nil {
		predicates = append(predicates, loginsession.LastLoginTimeLT(*i.LastLoginTimeLT))
	}
	if i.LastLoginTimeLTE != nil {
		predicates = append(predicates, loginsession.LastLoginTimeLTE(*i.LastLoginTimeLTE))
	}

	if i.HasOwner != nil {
		p := loginsession.HasOwner()
		if !*i.HasOwner {
			p = loginsession.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOwnerWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, loginsession.HasOwnerWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyLoginSessionWhereInput
	case 1:
		return predicates[0], nil
	default:
		return loginsession.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "email_credential" edge predicates.
	HasEmailCredential     *bool                        `json:"hasEmailCredential,omitempty"`
	HasEmailCredentialWith []*EmailCredentialWhereInput `json:"hasEmailCredentialWith,omitempty"`

	// "login_sessions" edge predicates.
	HasLoginSessions     *bool                     `json:"hasLoginSessions,omitempty"`
	HasLoginSessionsWith []*LoginSessionWhereInput `json:"hasLoginSessionsWith,omitempty"`

	// "public_profile" edge predicates.
	HasPublicProfile     *bool                          `json:"hasPublicProfile,omitempty"`
	HasPublicProfileWith []*UserPublicProfileWhereInput `json:"hasPublicProfileWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("entgenerated: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}

	if i.HasEmailCredential != nil {
		p := user.HasEmailCredential()
		if !*i.HasEmailCredential {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasEmailCredentialWith) > 0 {
		with := make([]predicate.EmailCredential, 0, len(i.HasEmailCredentialWith))
		for _, w := range i.HasEmailCredentialWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasEmailCredentialWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasEmailCredentialWith(with...))
	}
	if i.HasLoginSessions != nil {
		p := user.HasLoginSessions()
		if !*i.HasLoginSessions {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasLoginSessionsWith) > 0 {
		with := make([]predicate.LoginSession, 0, len(i.HasLoginSessionsWith))
		for _, w := range i.HasLoginSessionsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasLoginSessionsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasLoginSessionsWith(with...))
	}
	if i.HasPublicProfile != nil {
		p := user.HasPublicProfile()
		if !*i.HasPublicProfile {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasPublicProfileWith) > 0 {
		with := make([]predicate.UserPublicProfile, 0, len(i.HasPublicProfileWith))
		for _, w := range i.HasPublicProfileWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasPublicProfileWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasPublicProfileWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}

// UserPublicProfileWhereInput represents a where input for filtering UserPublicProfile queries.
type UserPublicProfileWhereInput struct {
	Predicates []predicate.UserPublicProfile  `json:"-"`
	Not        *UserPublicProfileWhereInput   `json:"not,omitempty"`
	Or         []*UserPublicProfileWhereInput `json:"or,omitempty"`
	And        []*UserPublicProfileWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "handle_name" field predicates.
	HandleName             *string  `json:"handleName,omitempty"`
	HandleNameNEQ          *string  `json:"handleNameNEQ,omitempty"`
	HandleNameIn           []string `json:"handleNameIn,omitempty"`
	HandleNameNotIn        []string `json:"handleNameNotIn,omitempty"`
	HandleNameGT           *string  `json:"handleNameGT,omitempty"`
	HandleNameGTE          *string  `json:"handleNameGTE,omitempty"`
	HandleNameLT           *string  `json:"handleNameLT,omitempty"`
	HandleNameLTE          *string  `json:"handleNameLTE,omitempty"`
	HandleNameContains     *string  `json:"handleNameContains,omitempty"`
	HandleNameHasPrefix    *string  `json:"handleNameHasPrefix,omitempty"`
	HandleNameHasSuffix    *string  `json:"handleNameHasSuffix,omitempty"`
	HandleNameEqualFold    *string  `json:"handleNameEqualFold,omitempty"`
	HandleNameContainsFold *string  `json:"handleNameContainsFold,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserPublicProfileWhereInput) AddPredicates(predicates ...predicate.UserPublicProfile) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserPublicProfileWhereInput filter on the UserPublicProfileQuery builder.
func (i *UserPublicProfileWhereInput) Filter(q *UserPublicProfileQuery) (*UserPublicProfileQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserPublicProfileWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserPublicProfileWhereInput is returned in case the UserPublicProfileWhereInput is empty.
var ErrEmptyUserPublicProfileWhereInput = errors.New("entgenerated: empty predicate UserPublicProfileWhereInput")

// P returns a predicate for filtering userpublicprofiles.
// An error is returned if the input is empty or invalid.
func (i *UserPublicProfileWhereInput) P() (predicate.UserPublicProfile, error) {
	var predicates []predicate.UserPublicProfile
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, userpublicprofile.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.UserPublicProfile, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, userpublicprofile.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.UserPublicProfile, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, userpublicprofile.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, userpublicprofile.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, userpublicprofile.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, userpublicprofile.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, userpublicprofile.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, userpublicprofile.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, userpublicprofile.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, userpublicprofile.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, userpublicprofile.IDLTE(*i.IDLTE))
	}
	if i.HandleName != nil {
		predicates = append(predicates, userpublicprofile.HandleNameEQ(*i.HandleName))
	}
	if i.HandleNameNEQ != nil {
		predicates = append(predicates, userpublicprofile.HandleNameNEQ(*i.HandleNameNEQ))
	}
	if len(i.HandleNameIn) > 0 {
		predicates = append(predicates, userpublicprofile.HandleNameIn(i.HandleNameIn...))
	}
	if len(i.HandleNameNotIn) > 0 {
		predicates = append(predicates, userpublicprofile.HandleNameNotIn(i.HandleNameNotIn...))
	}
	if i.HandleNameGT != nil {
		predicates = append(predicates, userpublicprofile.HandleNameGT(*i.HandleNameGT))
	}
	if i.HandleNameGTE != nil {
		predicates = append(predicates, userpublicprofile.HandleNameGTE(*i.HandleNameGTE))
	}
	if i.HandleNameLT != nil {
		predicates = append(predicates, userpublicprofile.HandleNameLT(*i.HandleNameLT))
	}
	if i.HandleNameLTE != nil {
		predicates = append(predicates, userpublicprofile.HandleNameLTE(*i.HandleNameLTE))
	}
	if i.HandleNameContains != nil {
		predicates = append(predicates, userpublicprofile.HandleNameContains(*i.HandleNameContains))
	}
	if i.HandleNameHasPrefix != nil {
		predicates = append(predicates, userpublicprofile.HandleNameHasPrefix(*i.HandleNameHasPrefix))
	}
	if i.HandleNameHasSuffix != nil {
		predicates = append(predicates, userpublicprofile.HandleNameHasSuffix(*i.HandleNameHasSuffix))
	}
	if i.HandleNameEqualFold != nil {
		predicates = append(predicates, userpublicprofile.HandleNameEqualFold(*i.HandleNameEqualFold))
	}
	if i.HandleNameContainsFold != nil {
		predicates = append(predicates, userpublicprofile.HandleNameContainsFold(*i.HandleNameContainsFold))
	}

	if i.HasOwner != nil {
		p := userpublicprofile.HasOwner()
		if !*i.HasOwner {
			p = userpublicprofile.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOwnerWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, userpublicprofile.HasOwnerWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserPublicProfileWhereInput
	case 1:
		return predicates[0], nil
	default:
		return userpublicprofile.And(predicates...), nil
	}
}
