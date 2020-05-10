use cursive::{Cursive, View};
use cursive::theme::{Color, PaletteColor, Theme};
use cursive::traits::{Nameable, Resizable, Scrollable};
use cursive::views::{Dialog, EditView, LinearLayout, ListView, ResizedView, TextView, ViewRef, ScrollView};
use cursive::view::ScrollStrategy;

fn main() {
    let mut cursive_root = Cursive::default();

    let mut message_pane = LinearLayout::vertical();

    let message_history = TextView::empty()
        .scrollable()
        .show_scrollbars(false)
        .scroll_strategy(ScrollStrategy::StickToBottom)
        .full_screen()
        .with_name("message_history");
    message_pane.add_child(message_history);

    message_pane.add_child(EditView::new()
        .on_submit(handle_message_submit)
        .full_width()
        .with_name("send_message"));

    cursive_root.set_theme(configure_theme(&cursive_root));
    cursive_root.screen_mut().add_transparent_layer(message_pane);
    cursive_root.run();
}

// TODO: There is some bug here where if you send messages super quick it'll hold the previous message state and just mutate it <x> characters.
// TODO: so if you send something like 00000, super quickly then send a 1, it'll do: 100000 for some reason.
fn handle_message_submit(cursive_root: &mut Cursive, message: &str) {
    cursive_root.call_on_name("message_history", |view: &mut ResizedView<ScrollView<TextView>> | {
        view.get_inner_mut().get_inner_mut().append(message.to_owned() + "\n");
        view.get_inner_mut().set_scroll_strategy(ScrollStrategy::StickToBottom);
    });

    // Reset the send message input once we have rendered the message.
    cursive_root.call_on_name("send_message", | view: &mut ResizedView<EditView> | {
        view.get_inner_mut().set_content("");
    });
}

fn configure_theme(cursive_root: &Cursive) -> Theme {
    let mut theme = cursive_root.current_theme().clone();
    theme.palette[PaletteColor::Background] = Color::TerminalDefault;

    return theme;
}