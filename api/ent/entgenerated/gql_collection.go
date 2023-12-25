// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/emailcredential"
	"api/ent/entgenerated/loginsession"
	"api/ent/entgenerated/userpublicprofile"
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ec *EmailCredentialQuery) CollectFields(ctx context.Context, satisfies ...string) (*EmailCredentialQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ec, nil
	}
	if err := ec.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ec, nil
}

func (ec *EmailCredentialQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(emailcredential.Columns))
		selectedFields = []string{emailcredential.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "owner":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: ec.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ec.withOwner = query
			if _, ok := fieldSeen[emailcredential.FieldOwnerID]; !ok {
				selectedFields = append(selectedFields, emailcredential.FieldOwnerID)
				fieldSeen[emailcredential.FieldOwnerID] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[emailcredential.FieldEmail]; !ok {
				selectedFields = append(selectedFields, emailcredential.FieldEmail)
				fieldSeen[emailcredential.FieldEmail] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ec.Select(selectedFields...)
	}
	return nil
}

type emailcredentialPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []EmailCredentialPaginateOption
}

func newEmailCredentialPaginateArgs(rv map[string]any) *emailcredentialPaginateArgs {
	args := &emailcredentialPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*EmailCredentialWhereInput); ok {
		args.opts = append(args.opts, WithEmailCredentialFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ls *LoginSessionQuery) CollectFields(ctx context.Context, satisfies ...string) (*LoginSessionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ls, nil
	}
	if err := ls.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ls, nil
}

func (ls *LoginSessionQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(loginsession.Columns))
		selectedFields = []string{loginsession.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "owner":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: ls.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ls.withOwner = query
			if _, ok := fieldSeen[loginsession.FieldOwnerID]; !ok {
				selectedFields = append(selectedFields, loginsession.FieldOwnerID)
				fieldSeen[loginsession.FieldOwnerID] = struct{}{}
			}
		case "lastLoginTime":
			if _, ok := fieldSeen[loginsession.FieldLastLoginTime]; !ok {
				selectedFields = append(selectedFields, loginsession.FieldLastLoginTime)
				fieldSeen[loginsession.FieldLastLoginTime] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ls.Select(selectedFields...)
	}
	return nil
}

type loginsessionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LoginSessionPaginateOption
}

func newLoginSessionPaginateArgs(rv map[string]any) *loginsessionPaginateArgs {
	args := &loginsessionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*LoginSessionWhereInput); ok {
		args.opts = append(args.opts, WithLoginSessionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "emailCredential":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&EmailCredentialClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.withEmailCredential = query
		case "loginSessions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&LoginSessionClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedLoginSessions(alias, func(wq *LoginSessionQuery) {
				*wq = *query
			})
		case "publicProfile":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserPublicProfileClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.withPublicProfile = query
		}
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (upp *UserPublicProfileQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserPublicProfileQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return upp, nil
	}
	if err := upp.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return upp, nil
}

func (upp *UserPublicProfileQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(userpublicprofile.Columns))
		selectedFields = []string{userpublicprofile.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "owner":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: upp.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			upp.withOwner = query
			if _, ok := fieldSeen[userpublicprofile.FieldOwnerID]; !ok {
				selectedFields = append(selectedFields, userpublicprofile.FieldOwnerID)
				fieldSeen[userpublicprofile.FieldOwnerID] = struct{}{}
			}
		case "handleName":
			if _, ok := fieldSeen[userpublicprofile.FieldHandleName]; !ok {
				selectedFields = append(selectedFields, userpublicprofile.FieldHandleName)
				fieldSeen[userpublicprofile.FieldHandleName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		upp.Select(selectedFields...)
	}
	return nil
}

type userpublicprofilePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPublicProfilePaginateOption
}

func newUserPublicProfilePaginateArgs(rv map[string]any) *userpublicprofilePaginateArgs {
	args := &userpublicprofilePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserPublicProfileWhereInput); ok {
		args.opts = append(args.opts, WithUserPublicProfileFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}